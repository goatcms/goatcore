package smtpmail

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
	"testing"
)

func base64TestExpectetTheSame(t *testing.T, s string) {
	var (
		buff        []byte
		resultBytes []byte
		result      string
		err         error
	)
	encoder := NewBase64Encoder(strings.NewReader(s))
	if buff, err = ioutil.ReadAll(encoder); err != nil {
		t.Error(err)
		return
	}
	if resultBytes, err = base64.StdEncoding.DecodeString(string(buff)); err != nil {
		t.Error(err)
		return
	}
	result = string(resultBytes)
	if result != s {
		t.Errorf("Expected '%v' and take '%v'", s, result)
	}
}

func TestBase64Converter(t *testing.T) {
	t.Parallel()
	base64TestExpectetTheSame(t, `OA`)
	base64TestExpectetTheSame(t, `Ala ma kota`)
	base64TestExpectetTheSame(t, `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse sit amet lorem lacus. Sed venenatis ligula in dapibus hendrerit. Aliquam sollicitudin fringilla felis, in imperdiet est laoreet et. Donec faucibus sollicitudin leo nec egestas. Sed dapibus posuere magna. In dapibus ante vitae convallis laoreet. Praesent dui nunc, feugiat ac ultricies ut, lacinia et diam. Donec et mi sit amet tortor malesuada porta. In et turpis augue. Nunc varius ullamcorper neque a convallis.

Ut dui purus, tempus ac auctor sit amet, tempus et est. Sed accumsan ex sit amet dolor lobortis rhoncus. Sed efficitur risus in purus molestie ultrices. In dapibus tellus et turpis euismod aliquet. Duis nec nisi nunc. Maecenas tincidunt lacinia tellus sit amet finibus. Nulla sit amet dictum dui. Phasellus et lorem tincidunt, tristique diam quis, placerat arcu. Integer elementum turpis et libero dignissim efficitur non in ante.

Aliquam pretium ante dolor, ut posuere turpis mattis sit amet. Pellentesque id libero ut odio tempor laoreet. Donec non facilisis tortor, in rhoncus urna. Integer bibendum sem non mattis lobortis. Fusce ut mauris at elit rutrum porttitor. Fusce dictum quam dolor, nec egestas mauris lacinia et. Etiam vehicula lectus vitae dui lacinia interdum. Nam ut imperdiet leo. Sed quis libero a turpis porttitor malesuada quis eget turpis. Donec tempus risus a risus pulvinar vestibulum. Ut fermentum lacus lacinia, fringilla est non, vehicula eros. Proin et sem eget nisl varius posuere.

Vivamus quis egestas tellus. Cras tempus viverra tellus, ac iaculis lectus scelerisque non. Etiam lacus sapien, semper ornare ligula in, laoreet facilisis velit. Integer ligula neque, dictum non faucibus quis, luctus non nisi. Quisque rutrum, odio ut malesuada semper, massa nunc malesuada augue, ut eleifend nibh ipsum in tortor. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Nunc blandit et ex ut elementum. Mauris maximus diam et magna faucibus, quis congue sapien fringilla. Cras viverra fermentum consectetur. In est mi, euismod placerat tincidunt in, luctus a tortor. Morbi vulputate tempus velit a volutpat. Maecenas faucibus eget nulla id volutpat. Vivamus egestas erat orci, viverra convallis metus tincidunt id. Vivamus quis diam quis magna congue egestas.

Aliquam diam enim, vestibulum non erat id, tincidunt dictum neque. Morbi mollis semper diam, quis ultricies nunc semper rutrum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In vitae eros eu quam varius gravida in non leo. Aliquam erat volutpat. Proin iaculis, mi a varius fringilla, nisl quam scelerisque odio, ut pharetra mi enim eget velit. Nam malesuada tellus vel enim ornare accumsan vel a arcu. Aenean mollis tellus risus, eget laoreet arcu interdum vitae. Integer dapibus purus et ultricies posuere. In placerat semper felis, commodo ultrices lacus sollicitudin a. Aliquam tincidunt rutrum felis.
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse sit amet lorem lacus. Sed venenatis ligula in dapibus hendrerit. Aliquam sollicitudin fringilla felis, in imperdiet est laoreet et. Donec faucibus sollicitudin leo nec egestas. Sed dapibus posuere magna. In dapibus ante vitae convallis laoreet. Praesent dui nunc, feugiat ac ultricies ut, lacinia et diam. Donec et mi sit amet tortor malesuada porta. In et turpis augue. Nunc varius ullamcorper neque a convallis.

Ut dui purus, tempus ac auctor sit amet, tempus et est. Sed accumsan ex sit amet dolor lobortis rhoncus. Sed efficitur risus in purus molestie ultrices. In dapibus tellus et turpis euismod aliquet. Duis nec nisi nunc. Maecenas tincidunt lacinia tellus sit amet finibus. Nulla sit amet dictum dui. Phasellus et lorem tincidunt, tristique diam quis, placerat arcu. Integer elementum turpis et libero dignissim efficitur non in ante.

Aliquam pretium ante dolor, ut posuere turpis mattis sit amet. Pellentesque id libero ut odio tempor laoreet. Donec non facilisis tortor, in rhoncus urna. Integer bibendum sem non mattis lobortis. Fusce ut mauris at elit rutrum porttitor. Fusce dictum quam dolor, nec egestas mauris lacinia et. Etiam vehicula lectus vitae dui lacinia interdum. Nam ut imperdiet leo. Sed quis libero a turpis porttitor malesuada quis eget turpis. Donec tempus risus a risus pulvinar vestibulum. Ut fermentum lacus lacinia, fringilla est non, vehicula eros. Proin et sem eget nisl varius posuere.

Vivamus quis egestas tellus. Cras tempus viverra tellus, ac iaculis lectus scelerisque non. Etiam lacus sapien, semper ornare ligula in, laoreet facilisis velit. Integer ligula neque, dictum non faucibus quis, luctus non nisi. Quisque rutrum, odio ut malesuada semper, massa nunc malesuada augue, ut eleifend nibh ipsum in tortor. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Nunc blandit et ex ut elementum. Mauris maximus diam et magna faucibus, quis congue sapien fringilla. Cras viverra fermentum consectetur. In est mi, euismod placerat tincidunt in, luctus a tortor. Morbi vulputate tempus velit a volutpat. Maecenas faucibus eget nulla id volutpat. Vivamus egestas erat orci, viverra convallis metus tincidunt id. Vivamus quis diam quis magna congue egestas.

Aliquam diam enim, vestibulum non erat id, tincidunt dictum neque. Morbi mollis semper diam, quis ultricies nunc semper rutrum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; In vitae eros eu quam varius gravida in non leo. Aliquam erat volutpat. Proin iaculis, mi a varius fringilla, nisl quam scelerisque odio, ut pharetra mi enim eget velit. Nam malesuada tellus vel enim ornare accumsan vel a arcu. Aenean mollis tellus risus, eget laoreet arcu interdum vitae. Integer dapibus purus et ultricies posuere. In placerat semper felis, commodo ultrices lacus sollicitudin a. Aliquam tincidunt rutrum felis.`)
}
