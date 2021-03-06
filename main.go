package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/go-echarts/go-echarts/v2/charts"
    "github.com/go-echarts/go-echarts/v2/opts"
    "github.com/go-echarts/go-echarts/v2/types"
    "fmt"
    "os"
    "encoding/json"
    "path/filepath"
    "log"
    "html/template"
    "ab_log_plotter/models"
    "ab_log_plotter/configer"
    "time"
   _ "github.com/lib/pq"
)



// структура для реле, заполняем состояниями
type LightRelays struct{
    R_id int
    R_ip string
    R_state int
    R_update string
}



//config reader
func Config_reader(cfg_file string) configer.Configuration {

	file, err := os.Open(cfg_file)
	if err != nil {
		fmt.Println("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := configer.Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("can't decode config JSON: ", err)
	}

	return Config
}





func AddRoutes(r *mux.Router) {
	r.HandleFunc("/get/{category}/", getWriter).Methods("GET")
//	r.HandleFunc("/get/{category}/", getWriter).Methods("GET")
//	r.HandleFunc("/get/relays/", showRelays).Methods("GET")
}



// generate random data for line chart
func generateLineItems(query_val string) []opts.LineData {

	rows, err := models.Db.Query(query_val)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	items := make([]opts.LineData, 0)
	for rows.Next() {
		var ls float64
		if err := rows.Scan(&ls); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
                items = append(items, opts.LineData{Value: ls})
	}
	return items
}



// generate random data for line chart
func generatetimeline(query_date string) []time.Time {

	rows, err := models.Db.Query(query_date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	items := make([]time.Time, 0)
	for rows.Next() {
		var ls time.Time
		if err := rows.Scan(&ls); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
                items = append(items, ls)
	}
	return items
}



func getWriter(w http.ResponseWriter, r *http.Request) {
        //категория : температура, освещенность , ветер и тд
        vars := mux.Vars(r)
        cat := vars["category"]
        temp_query_date :=  "SELECT w_date FROM weather WHERE w_date BETWEEN NOW()- INTERVAL '72 HOURS' AND NOW()  ORDER BY w_date asc"
        temp_query_val  :=  "SELECT temp_val FROM weather WHERE w_date BETWEEN NOW() - INTERVAL '72 HOURS' AND NOW()  order BY w_date asc"
        temp_title := "График датчика температуры д.Новая Слободка"

        press_query_date :=  "SELECT TO_TIMESTAMP(TO_CHAR(p_date,'HH24:MI:SS DD-MM-YYYY'),'HH24:MI:SS DD-MM-YYYY') as p_date FROM pressure WHERE p_date BETWEEN NOW()- INTERVAL '144 HOURS' AND NOW()  ORDER BY p_date asc"
        press_query_val  :=  "SELECT p_val-700.00  FROM pressure WHERE p_date BETWEEN NOW() - INTERVAL '144 HOURS' AND NOW()  order BY p_date asc"
        press_title := "График датчика давления д.Новая Слободка"

        light_query_date := "SELECT light_date FROM light WHERE  light_date BETWEEN NOW()- INTERVAL '72 HOURS' AND NOW()  ORDER BY light_date asc"
        light_query_val  := "SELECT light_val FROM light WHERE light_date BETWEEN NOW() - INTERVAL '72 HOURS' AND NOW()  order BY light_date asc"
        light_title := "График датчика освещения  д.Новая Слободка"

        temprasp_query_date :=  "SELECT r_date FROM rasp_weather WHERE r_date BETWEEN NOW()- INTERVAL '72 HOURS' AND NOW()  ORDER BY r_date asc"
        temprasp_query_val  :=  "SELECT temp_val FROM rasp_weather WHERE r_date BETWEEN NOW() - INTERVAL '72 HOURS' AND NOW()  order BY r_date asc"
        temprasp_title := "График датчика температуры Raspberry  д.Новая Слободка"



        if cat == "temp"{
		// create a new line instance
		line := charts.NewLine()
		// set some global options like Title/Legend/ToolTip or anything else
		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
			charts.WithTitleOpts(opts.Title{
				Title:    temp_title,
				Subtitle: getCurrentTemp(),
			}))
		// Put data into instance
/*		line.SetXAxis( generatetimeline(temp_query_date)).
			AddSeries("Category A", generateLineItems(temp_query_val)).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
*/
		line.SetXAxis( generatetimeline(temp_query_date)).
			AddSeries("Category A", generateLineItems(temp_query_val)).
			SetSeriesOptions(
                                 charts.WithLabelOpts(opts.Label{
                                         Show: true,
                                 }),
                                 charts.WithAreaStyleOpts(opts.AreaStyle{
                                         Opacity:0.2,
                                 }),
                                 charts.WithLineChartOpts(opts.LineChart{
                                         Smooth: true,
                                 }),
                        )
		line.Render(w)
	}else if cat == "temprasp"{
		// create a new line instance
		line := charts.NewLine()
		// set some global options like Title/Legend/ToolTip or anything else
		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
			charts.WithTitleOpts(opts.Title{
				Title:    temprasp_title,
				Subtitle:  "--",
			}))
		// Put data into instance
/*		line.SetXAxis( generatetimeline(temp_query_date)).
			AddSeries("Category A", generateLineItems(temp_query_val)).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
*/
		line.SetXAxis( generatetimeline(temprasp_query_date)).
			AddSeries("Category A", generateLineItems(temprasp_query_val)).
			SetSeriesOptions(
                                 charts.WithLabelOpts(opts.Label{
                                         Show: true,
                                 }),
                                 charts.WithAreaStyleOpts(opts.AreaStyle{
                                         Opacity:0.2,
                                 }),
                                 charts.WithLineChartOpts(opts.LineChart{
                                         Smooth: true,
                                 }),
                        )
		line.Render(w)
	
          
        }else  if cat =="light"{

		// create a new line instance
		line := charts.NewLine()
		// set some global options like Title/Legend/ToolTip or anything else
		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
			charts.WithTitleOpts(opts.Title{
				Title:    light_title,
				Subtitle: "--",
			}))
		// Put data into instance
		line.SetXAxis( generatetimeline(light_query_date)).
			AddSeries("Category A", generateLineItems(light_query_val)).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
		line.Render(w)
        }else if cat == "relays"{
		getstate:= getRelaystate()
                fmt.Println(getstate)
		tmpl, err := template.ParseFiles("templates/index.html")
		if  err != nil {
                   log.Println(err)
		}
		err=tmpl.Execute(w, getstate)
		if  err != nil {
			log.Println(err)
		}
        }else  if cat == "press"{
		// create a new line instance
		line := charts.NewLine()
		// set some global options like Title/Legend/ToolTip or anything else
		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
			charts.WithTitleOpts(opts.Title{
				Title:    press_title,
				Subtitle: "",
			}),
                charts.WithXAxisOpts(opts.XAxis{
			Name: "Время",
			//Type: "category",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Давление +700мм.рт.ст",
			//Type: "value",
		}),

                )

		line.SetXAxis( generatetimeline(press_query_date)).
			AddSeries("Category A", generateLineItems(press_query_val)).
			SetSeriesOptions(
                                 charts.WithLabelOpts(opts.Label{
                                         Show: true,
                                 }),
                                 charts.WithAreaStyleOpts(opts.AreaStyle{
                                         Opacity:0.2,
                                 }),
                                 charts.WithLineChartOpts(opts.LineChart{
                                         Smooth: true,
                                 }),
                        )
		line.Render(w)
	}

}




func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }
/****************************** Relays *******************************************/
/*
//шаблонизатор отображения странички с состоянием реле
func showRelays(w http.ResponseWriter, r *http.Request) {
        getstate:= getRelaystate()
        tmpl, err := template.ParseFiles("templates/index.html")
        if  err != nil {
            fmt.Println(err)
        }
        err=tmpl.Execute(w, getstate)
        if  err != nil {
            fmt.Println(err)
        }

}
*/


func getCurrentTemp()(curTemp string) {
        var currentTemp float64
        err := models.Db.QueryRow("SELECT temp_val FROM weather  order BY w_date desc limit 1").Scan(&currentTemp)
        if err != nil {
             log.Fatal(err)
             curTemp = "N/A"
        }
        curTemp = fmt.Sprintf("%f",currentTemp)
        return curTemp
}



//получаем состояние реле
func getRelaystate() []LightRelays {

        rows, err := models.Db.Query("SELECT r_id, r_ip, r_state, to_char(r_last_update, 'HH24:MI:SS DD Mon YYYY')  FROM relays WHERE r_type=1  ORDER BY r_id asc")
        if err != nil {
             log.Fatal(err)
        }
        defer rows.Close()

        lightrelays := []LightRelays{}

        for rows.Next(){
             lr := LightRelays{}
             err := rows.Scan(&lr.R_id, &lr.R_ip, &lr.R_state,&lr.R_update)
             if err != nil{
                  log.Println(err)
                  continue
             }
             lightrelays = append(lightrelays, lr)
        }


    return lightrelays

}




/*******************************************************************************/


func main() {

     version := "0.0.2"
     fmt.Println("ab-log plotter version:"+version)
//************************* read config ******************************************//
     dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
     if err != nil {
            log.Fatal(err)
     }

     //log.Println(dir)
     cfg := Config_reader(filepath.Join(dir,"ploter.conf"))
     models.Initdb(cfg)

//*********************** parse config **********************************//
   //logging
   log_dir := "./log"
   if _, err := os.Stat(log_dir); os.IsNotExist(err) {
		os.Mkdir(log_dir, 0644)
   }
   file, err := os.OpenFile(filepath.Join(log_dir,cfg.Log_file_name), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
   if err != nil {
		log.Fatal(err)
   }
   defer file.Close()
   log.SetOutput(file)
   log.Println("Logging to a file plotter!")

   router := mux.NewRouter()
   AddRoutes(router)
   log.Fatal(http.ListenAndServe(":80", router))

}
