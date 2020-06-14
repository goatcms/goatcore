package sshsb

import (
	"bytes"
	"io"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"golang.org/x/crypto/ssh"
)

// SSHSandbox is termal sandbox
type SSHSandbox struct {
	username   string
	host       string
	cwd        string
	entrypoint string
	deps       deps
}

// NewSSHSandbox create a SSHSandbox instance
func NewSSHSandbox(sshParam, entrypoint string, deps deps) (ins pipservices.Sandbox, err error) {
	var (
		username string
		host     string
		index    int
	)
	sshParam = strings.Trim(sshParam, " \t\n")
	if index = strings.Index(sshParam, "@"); index == -1 {
		return nil, goaterr.Errorf("Insert ssh connection host like ssh:username@hostname. Username and host name (or IP) are required.")
	}
	if username = strings.Trim(sshParam[:index], " \t\n"); username == "" {
		return nil, goaterr.Errorf("SSH Sandbox: Username can not be empty. Insert sandbox like ssh:username@hostname")
	}
	if host = strings.Trim(sshParam[index+1:], " \t\n"); host == "" {
		return nil, goaterr.Errorf("SSH Sandbox: Host can not be empty. Insert sandbox like ssh:username@hostname")
	}
	return &SSHSandbox{
		username:   username,
		host:       host,
		entrypoint: entrypoint,
		deps:       deps,
	}, nil
}

// Run run code in sandbox
func (sandbox *SSHSandbox) Run(ctx app.IOContext) (err error) {
	var (
		cio        = ctx.IO()
		envs       commservices.Environments
		initReader io.Reader
		sshKey     ssh.Signer
		sshClient  *ssh.Client
		sshSession *ssh.Session
		buf        = &bytes.Buffer{}
		bufOutput  = gio.NewOutput(buf)
	)
	if envs, err = sandbox.deps.EnvironmentsUnit.Envs(ctx.Scope()); err != nil {
		return err
	}
	if initReader, err = sandbox.initSequence(envs); err != nil {
		return err
	}
	sshCert := envs.SSHCert()
	if sshCert.Public == "" {
		return goaterr.Errorf("SSHCert: Public key is required")
	}
	if sshCert.Secret == "" {
		return goaterr.Errorf("SSHCert: Secret key is required")
	}
	if sshKey, err = ssh.ParsePrivateKey([]byte(sshCert.Secret)); err != nil {
		return goaterr.Wrap(err, "SSH Sandbox: ParsePrivateKey error")
	}
	if sshClient, err = ssh.Dial("tcp", sandbox.host, &ssh.ClientConfig{
		User: sandbox.username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(sshKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}); err != nil {
		return goaterr.Wrapf("SSH Sandbox: Connect to %s@%s error", err, sandbox.username, sandbox.host)
	}
	defer sshClient.Close()
	if sshSession, err = sshClient.NewSession(); err != nil {
		return goaterr.Wrapf("SSH Sandbox: %s@%s session error", err, sandbox.username, sandbox.host)
	}
	defer sshSession.Close()
	repeatIO := gio.NewRepeatIO(gio.IOParams{
		In:  cio.In(),
		Out: bufOutput,
		Err: bufOutput,
		CWD: cio.CWD(),
	})
	sshSession.Stdin = io.MultiReader(initReader, repeatIO.In())
	sshSession.Stdout = io.MultiWriter(repeatIO.Out(), cio.Out())
	sshSession.Stderr = io.MultiWriter(repeatIO.Err(), cio.Err())
	if err = sshSession.Shell(); err != nil {
		return goaterr.Wrapf("SSH Sandbox: %s@%s Execution error", err, sandbox.username, sandbox.host, buf.String())
	}
	if err = sshSession.Wait(); err != nil {
		return goaterr.Wrapf("SSH Sandbox: %s@%s Execution error\n%s", err, sandbox.username, sandbox.host, buf.String())
	}
	return nil
}

func (sandbox *SSHSandbox) initSequence(envs commservices.Environments) (reader io.Reader, err error) {
	var (
		initCode = "\nset -e\nset +x\n"
		eofTag   = "EOF" + varutil.RandString(10, varutil.UpperAlphaBytes)
	)
	for key, value := range envs.All() {
		initCode += key + "=$(cat <<" + eofTag + "\n" + value + "\n" + eofTag + "\n)\n"
		initCode += "export " + key + "\n"
	}
	initCode += sandbox.entrypoint + "\n"
	return strings.NewReader(initCode), nil
}
