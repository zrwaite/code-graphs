import cron from "node-cron";
import drawGraph from "../drawGraph";
const cronjobs = () => {
	drawGraph();
	cron.schedule('0 1 * * *', () => {
		drawGraph();
	});
}
export default cronjobs;