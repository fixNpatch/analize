function CreateColumns(data, param) { // data = alphabet, param = which alphabet (rus/eng)
	for(let i in data) {
		let tablename = param + "_datatable_part" // + что-то там с i
		$$(tablename).define({
			columns:[]
		})
	}
}