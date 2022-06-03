# Weather CLI

Weather CLI is a command line tool which helps you to learn current forecast information of city that you specify.

On this project I used [Cobra](https://github.com/spf13/cobra), Go based framework.

---
## Installation

```bash
git clone https://github.com/emrullahcirit/weather-cli
cd path/to/weather-cli
```

After you get into project directory you need to get an API Key from [Open Weather Map](https://openweathermap.org/appid) then you need to creat a .env file and paste you api key into it.

```bash
go build .
./weather-cli forecast --city <<city name here>>
```

---

## ScreenShots

<img width="325" alt="Screen Shot 2022-06-03 at 23 54 37" src="https://user-images.githubusercontent.com/55560241/171953577-ee147b90-25d9-42c4-af15-e2c22e8576f9.png">


<img width="627" alt="Screen Shot 2022-06-04 at 00 12 06" src="https://user-images.githubusercontent.com/55560241/171953616-3480a3fa-a79f-4dd5-aec2-0e5485196dea.png">



