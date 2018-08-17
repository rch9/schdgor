Чинуть про range


func (sc *Scheduler) Add(jobs ...Job) {
	for i, j := range jobs {
		j.status = StatReady
		j.stop = make(chan struct{}, 1) // TODO: buffered or unbuffered
		sc.JobsPool[j.Name] = &jobs[i]
		// fmt.Println(sc.JobsPool)
		fmt.Println("test", &jobs)
		fmt.Println(&j)
	}
	fmt.Println("jjj", jobs)
	fmt.Println("fuck", sc.JobsPool)
	// fmt.Println("fuck", sc.JobsPool["Job-1"])
	// fmt.Println("fuck", sc.JobsPool["Job-2"])
}
