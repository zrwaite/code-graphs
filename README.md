# CodeGraphs

### API service to display code data as static SVGs

See more info and authorize your account [here](https://graphs.insomnizac.com)

## `/api/wakatime/:username`:
### Pi graph that updates every hour with a cronjob pulling data from the WakaTime API
- Query Parameters:
	- **`ignore=html,css,scss`**: removes specific languages from the graph
	- **`addUsername=true`**: displays your username above the pi graph
	- **`removeDefaultIgnore=true`**: stops ignoring certain default languages (JSON, YML, Text, Markdown, etc)


<a href="https://graphs.insomnizac.com/api/wakatime/Insomnizac" target="_blank">
<img src="https://graphs.insomnizac.com/api/wakatime/Insomnizac" />
</a>

## `/api/streak/:username`:
### Static GitHub streak graph with dark mode styles

<a href="https://graphs.insomnizac.com/api/streak/zrwaite" target="_blank">
<img src="https://graphs.insomnizac.com/api/streak/zrwaite" />
</a>
