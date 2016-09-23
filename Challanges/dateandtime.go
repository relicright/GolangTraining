package main

import(
  "fmt"
  "time"
)

func main()  {
  //TIME.DATE ALLOWS YOU TO RETURN A SPECIFIC DATE THROUGH THE TIME PACKAGE
  //            YEAR       MONTH     DAY  HR  MIN SEC NANO SEC  TIME-ZONE
  t := time.Date(2009, time.November, 10, 23, 00, 00, 00, time.UTC)
  fmt.Printf("Go launched at: %s\n", t)

  // NOW RETURNS THE DATE AND TIME FROM THE SYSTEM IT IS RUNNING ON
  now := time.Now()
  fmt.Printf("The time now is: %s\n", now)

  //INDIVIDUAL TIMES AND DATES CAN BE RETURNED THROUGH THE TIME PACKAGE
  fmt.Println("the month is:", t.Month())
  fmt.Println("the day is:", t.Day())
  fmt.Println("the weekday is:", t.Weekday())

  // YOU CAN ADD DATES TO A CURRENT 'TIME.DATE' USING .AddDate
  //                year-month-day
  tomorrow := t.AddDate(0, 0, 1)
  fmt.Printf("Tomorrow is %v %v %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

  //YOU CAN SET THE FORMAT THE DATE DISPLAYS BY CRAFTING YOUR OWN FORMAT STRINGS, THEN USE .FORMAT()
  // the format is ordered in a specific way that can be found in golang.org/pkg/time/
  // monday = 0/ January = 1/ day = 2 / hour = 3 / min = 4 / second = 5/ year = 6
  longFormat := "Mon Jan 2 15:04:05 MST 2006"
  fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

  // ex: 2
  shortFormat := "1/2/06"
  fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
}
