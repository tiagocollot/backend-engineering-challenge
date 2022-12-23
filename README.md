# Tiago Collot - Backend Engineering Challenge

---
## Welcome to my solution 🖖

### Getting started

*To start the program run in the terminal:*

`go run . --input_file=data.json --window_size=10`



### Challenge Solution

*The moving average helps to level the duration data over a specified period by creating a constantly updated average delivery time.*

 *I have built a simple command line application that parses a stream of events and produces an aggregated output for a specified time period. In this case, I was interested in calculating, for every minute, a moving average of the translation delivery time for the last 10 minutes with the following input data:*

 ```json
 {"timestamp": "2018-12-26 18:11:08.509654","translation_id": "5aa5b2f39f7254a75aa5","source_language": "en","target_language": "fr","client_name": "easyjet","event_name": "translation_delivered","nr_words": 30, "duration": 20}
{"timestamp": "2018-12-26 18:15:19.903159","translation_id": "5aa5b2f39f7254a75aa4","source_language": "en","target_language": "fr","client_name": "easyjet","event_name": "translation_delivered","nr_words": 30, "duration": 31}
{"timestamp": "2018-12-26 18:23:19.903159","translation_id": "5aa5b2f39f7254a75bb33","source_language": "en","target_language": "fr","client_name": "booking","event_name": "translation_delivered","nr_words": 100, "duration": 54}
```
*And the corresponding output, giving the calculated moving average delivery time:*

 ```json
{"date":"2018-12-26 18:11:00","average_delivery_time":0}
{"date":"2018-12-26 18:12:00","average_delivery_time":20}
{"date":"2018-12-26 18:13:00","average_delivery_time":20}
{"date":"2018-12-26 18:14:00","average_delivery_time":20}
{"date":"2018-12-26 18:15:00","average_delivery_time":20}
{"date":"2018-12-26 18:16:00","average_delivery_time":25.5}
{"date":"2018-12-26 18:17:00","average_delivery_time":25.5}
{"date":"2018-12-26 18:18:00","average_delivery_time":25.5}
{"date":"2018-12-26 18:19:00","average_delivery_time":25.5}
{"date":"2018-12-26 18:20:00","average_delivery_time":25.5}
{"date":"2018-12-26 18:21:00","average_delivery_time":25.5}
{"date":"2018-12-26 18:22:00","average_delivery_time":31}
{"date":"2018-12-26 18:23:00","average_delivery_time":31}
{"date":"2018-12-26 18:24:00","average_delivery_time":42.5}
```

### Note:
*Thank you for this interesting challenge!*
