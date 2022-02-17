import path from "path";
import fs from "fs";
import axios from "axios";
import env from "dotenv";
env.config();

const swap = (arr: any[], i1:number, i2:number) => {
	let temp = arr[i1];
	arr[i1]=arr[i2];
	arr[i2]=temp;
}
const bubbleSortLanguages = (arr: any[]) => {
	for (let i=0; i<arr.length-1; i++){
		for (let j=0; j<arr.length-i-1; j++){
			if (arr[j].time<arr[j+1].time) swap(arr, j, j+1)
		}
	}
}
const getCodeData = async (token: string) => {
	let apiLink = "https://wakatime.com/api/v1/users/current/summaries?timeout=15&writes_only=true";
	let startDate = "2021/10/12"
	let date = new Date().toLocaleDateString().toString();
	apiLink += "&start="+startDate+"&end="+date;
	let headers = { headers: {'Host': 'wakatime.com', 'Authorization': 'Bearer '+token}};
	try{
		const codeData: any = await axios.get(apiLink,headers);
		return codeData.data;
	} catch(e:any){
		console.log("Error",e);
		return false;
	}
}

const writeLanguages = async () => {
	let ignoreLanguages = ["JSON", "Docker", "Markdown", "Other", "INI", "Text", "XML", "YAML", "Bash", "Git Config", "Objective-C", "TOML", "Apache Config", "GitIgnore file", "Shell Script"];	
	const folderPath = path.join(__dirname, "../json/");
	let fileName = "languages.json";
	let filePath = folderPath+fileName;
	let json = await getCodeData(`${process.env.WAKATIME_TOKEN}`)
	let languages: any[] = [];
	let totalTime: number = 0;
	json.data.forEach((data:any) => {
		data.languages.forEach((language:any)=>{
			let name = language.name;
			let time = language.total_seconds;
			let found = false;
			for (let i=0; i<languages.length; i++){
				if (name != languages[i].name) continue;
				languages[i].time += time;
				totalTime += time;
				found = true; 
				break;
			} 
			if (!found) {
				if (!ignoreLanguages.includes(name)){
					let colour;
					switch (name) {
						case "TypeScript": colour = "#0099ff"; break;
						case "JavaScript": colour = "#ecec13"; break;
						case "C": colour = "#666666"; break;
						case "C#": colour = "#9332bf"; break;
						case "JSON": colour = "#339933"; break;
						case "PHP": colour = "#9999ff"; break;
						case "Python": colour = "#0066cc"; break;
						case "HTML": colour = "#ff471a"; break;
						case "Docker": colour = "#1aa3ff"; break;
						case "SQL": colour = "#e6b800"; break;
						case "Java": colour = "#e60000"; break;
						case "Dart": colour = "rgb(23, 174, 255)"; break;
						case "SCSS": colour = "rgb(201, 85, 146)"; break;
						case "CSS": colour = "rgb(28, 49, 220)"; break;
						case "Rust": colour = "#ff5c33"; break;
						case "Racket": colour = "rgb(100, 13, 20)"; break;
						case "Markdown": colour = "#333333"; break;
						case "C++": colour = "rgb(83, 136, 200)"; break;
						case "VHDL": colour = "grey"; break;
						case "Go": colour = "rgb(20, 156, 206)"; break;
						default: colour = "white"; break;
					}
					languages.push({name:name, colour: colour, time:time, percent:0});
					totalTime += time;
				}
			}
			found = false;
		});
	})
	bubbleSortLanguages(languages);
	for (let i=0; i<languages.length; i++){
		languages[i].percent = (Math.round((languages[i].time/totalTime)*1000)/10);
		languages[i].time = Math.round(languages[i].time)
		if (languages[i].percent <= 0.1) languages.splice(i,languages.length-1);
	}	
	totalTime = Math.round(totalTime);
	fs.writeFileSync(filePath, JSON.stringify({time: totalTime, languages: languages}));
}

export default writeLanguages;