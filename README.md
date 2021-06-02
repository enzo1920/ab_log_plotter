# ab_log_plotter

## Построитель графиков из бд postgres
Программа предназначена для отображения значений датчиков, получаемых с контроллера [ab-log mega-d 2561](https://ab-log.ru/smart-house/ethernet/megad-2561)
К основному модулю программы(main.go) подключаются:
- модуль для инициализации подключения к бд
- модуль для работы с конфигурационном файлом(файл в формате json)

## Пути:
/light/ (отображение графика освещенности за 72 часа)

/temp/ (отображение графика температуры за 72 часа)

/relays/ (отображает состояние реле (вкл/выкл))

## Запуск
Запускается как web-server на 80 порту.
Используя screen:
```
screen -S plotter_session
./ab_log_grapher
```
## Log
Записывается в папку ./log

## Make
Для компиляции под raspberry pi используется комманда:
```
make rasp 
```
## Автостарт сервера на RPi
- Создать файл  .service в systemd directory:
```
sudo nano /lib/systemd/system/ablogplolt.service
```
- Содержание файла:
```
[Unit]
Description=ab-log plotter
After=multi-user.target

[Service]
WorkingDirectory=/home/ab_log_plotter/bin
ExecStart=/home/ab_log_plotter/bin/ab_log_grapher

[Install]
WantedBy=multi-user.target
```
- перегружаем systemd
```
sudo systemctl daemon-reload
```
- включаем
```
sudo systemctl enable ablogplolt.service
```
