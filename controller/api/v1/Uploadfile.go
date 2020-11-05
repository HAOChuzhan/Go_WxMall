package v1

/*
func Tolead(c *gin.Context) {
	appG := app.Gin{C: c}

	files, _ := c.FormFile("file")
	dst := path.Join("../upload", files.Filename)

	err := c.SaveUploadedFile(files, dst)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.InternalServerError, nil)
	}
	xlsx, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows, _ := xlsx.GetRows("Sheet" + "1")

	for key, row := range rows {
		if key > 0 {
			for _, colcell := range row {
				fmt.Println(colcell)
			}
		}
	}
	radio := banner.Radio{
		Title : row[0],
		JumpType: row[1]
	}
	models.DB.

}*/
