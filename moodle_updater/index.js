require('log-timestamp')(function() { return '[' + new Date() + '] %s'; });
const methods = require('./methods');
const schedule = require('node-schedule');
const db = require('./db');

function startScheduler() {
    if (process.env.NO_SCHEDULER == true) return;

    const UPDATE_CRONTAB = process.env.UPDATE_CRONTAB;
    console.log(`Starting scheduler with crontab "${UPDATE_CRONTAB}"`);
    schedule.scheduleJob(UPDATE_CRONTAB, async function () {
        await methods.updateAllAssignmentsWithSend();
    });
} 

if (process.env.SYNC_TABLES == true) {
    console.log('Syncing tables...');
    db.Assignment.sync();
} else {
    if (process.env.UPDATE_ON_START == true) {
        methods.updateAllAssignmentsWithSend().then(() => startScheduler());
    } else {
        startScheduler();
    }   
}
