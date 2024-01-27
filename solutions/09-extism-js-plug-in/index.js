function hello() {
	let data = Host.inputString()

	let text = "💛 Hello " + data

    console.log(text)

	let result = {
		message: text,
        from: "Extism"
	}

	Host.outputString(JSON.stringify(result))
}

module.exports = { hello }