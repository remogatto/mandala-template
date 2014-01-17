# What's that?

<tt>mandala-template</tt> generates a scaffolding to quickly get you started for writing a [Mandala](https://github.com/remogatto/mandala) application.

# Usage

<pre>
go get github.com/remogatto/mandala-template
mandala-template myapp
cd myapp
# Edit app.json for customization
gotask init
gotask run xorg # or
gotask run android
</pre>

# Black-box testing

<tt>mandala-template</tt> generates a scaffolding for a black-box test. To run the test:

<pre>
cd test
gotask test xorg # or
gotask test android
</pre>

# LICENSE

See [LICENSE](LICENSE).

