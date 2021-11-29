import cronjobs from "./cronjobs/cronjobs";
import env from "dotenv";
import app from "./server";
env.config();
const port = process.env.PORT || 2000;

app.listen(port, () => {
	console.log(`listening on port ${port}`);
	cronjobs();
});