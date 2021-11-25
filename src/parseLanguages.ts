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
				languages.push({name:name, time:time, percent:0});
				totalTime += time;
			}
			found = false;
		});
	})
	bubbleSortLanguages(languages);
	for (let i=0; i<languages.length; i++){
		languages[i].percent = (Math.round((languages[i].time/totalTime)*1000)/10);
		languages[i].time = Math.round(languages[i].time)
		if (languages[i].percent == 0) languages.splice(i,languages.length-1);
	}	
	totalTime = Math.round(totalTime);
	fs.writeFileSync(filePath, JSON.stringify({time: totalTime, languages: languages}));
}

export default writeLanguages;