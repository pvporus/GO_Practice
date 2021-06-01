package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstNAme: "raj",
			LastName:  "kulkarni",
			Grades: []Grade{
				{
					Title: "title1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "title2",
					Type:  GradeHomework,
					Score: 75,
				},
				{
					Title: "title3",
					Type:  GradeTest,
					Score: 65,
				},
			},
		},
		{
			ID:        2,
			FirstNAme: "mohan",
			LastName:  "kumar",
			Grades: []Grade{
				{
					Title: "title11",
					Type:  GradeQuiz,
					Score: 95,
				},
				{
					Title: "title22",
					Type:  GradeHomework,
					Score: 55,
				},
				{
					Title: "title32",
					Type:  GradeTest,
					Score: 87,
				},
			},
		},
	}
}
