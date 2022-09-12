package wakatime

// onst createLegend = async (ctx: any, languages: any[], start: number, width:number, height:number) => {
// 	// ctx.fillStyle = "lightgrey";
// 	// ctx.fillRect(start+ width/20, height/40, width-width/10, height-height/8);
// 	ctx.textBaseline = "top";
// 	ctx.textAlign = 'center';

// 	ctx.font = 'bold 45pt Menlo';
// 	ctx.fillStyle = "white";
// 	ctx.fillText("Top Languages", start+width/2, height/31);
// 	ctx.fillText("(By time spent)", start+width/2, 2*height/21);
// 	ctx.font = 'bold 18pt Menlo';
// 	ctx.fillText("Since October 2021*", start+width/2, 3.5*height/21);

// 	ctx.font = 'bold 35pt Menlo';
// 	ctx.fillStyle = "black";
// 	ctx.textAlign = 'left';
// 	for (let i=0; i<languages.length; i++) {
// 		let language = languages[i];
// 		ctx.fillStyle = language.colour;
// 		ctx.fillRect(start+width/11, (i+4.2)*height/19, height/20, height/35);

// 		ctx.shadowBlur=0;
// 		ctx.fillStyle=language.colour;
// 		ctx.fillText(language.name, start+2*width/10, (i+3.9)*height/19);

// 		ctx.shadowBlur=0;

// 		ctx.fillStyle = language.colour;
// 		ctx.fillText(language.percent.toString() + "%", start+7*width/10, (i+3.9)*height/19);
// 	}
// 	ctx.font = 'bold 24pt Menlo';
// 	ctx.fillStyle = "white";
// 	ctx.textAlign = 'center';
// 	let date = new Date().toLocaleDateString().toString();
// 	ctx.fillText("Charts coded by Zac in nodejs", start+width/2, 18.2*height/20);
// 	ctx.fillText("Updated on: "+date, start+width/2, 18.9*height/20);
// }
