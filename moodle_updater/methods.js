const moodle = require('./moodle');
const db = require('./db');
const request_promise = require('request-promise');

const SIGNAL_HOST = process.env.SIGNAL_HOST;
const SIGNAL_NUMBER = process.env.SIGNAL_NUMBER;


async function updateAssignmentsWithSend(token) {
    let a = await db.Assignment.findOne({
        where: {
            moodleToken: token
        }
    });

    if (!a){
        console.log(`token "${token}" not found.`);
        return false;
    } 
    

    let newRawAssignments = await moodle.getRawAssignmentsByToken(token);

    if (moodle.checkResponseForInvalidToken(newRawAssignments)) {
        await stopServiceForAssignment(a, 'Moodle-Anmeldedaten sind nicht (mehr) gültig. Bitte melde dich neu an.');
        return
    }
    
    let newAssignments = moodle.getAssignmentIdsByRawAssignments(newRawAssignments);

    let difference;
    difference = a.assignments ? newAssignments.filter(x => !a.assignments.includes(x)) : newAssignments;

    if (difference.length){
        // Send Signal notification
        if (a.assignments.length > 0) await sendSignalMessage(a.phoneNumber, textFromAssignmentIds(difference, newRawAssignments));

        a.assignments = newAssignments;
        await a.save();
    }
    return difference.length;
}


async function sendSignalMessage(to_number, message) {
    let options = {
        method: 'POST',
        uri: SIGNAL_HOST + '/v2/send',
        body: {
            message: message,
            number: SIGNAL_NUMBER,
            recipients: [to_number]
        },
        json: true
    };

    await request_promise(options).catch((e) => {
        console.error('Error occured when requesting to Signal Server:');
        console.error(e);
    });
}

async function updateAllAssignmentsWithSend() {
    console.log('Starting.')
    console.log('Getting tokens from DB...')
    let tokens = await getAllAccountTokens();
    console.log(`${tokens.length} token(s) found.`)
    if (!tokens) {
        console.log('No tokens, breaking up.');
        return;
    }
    for (let t in tokens) {
        console.log(`Updating user with moodleToken..."${tokens[t]}"`);
        let res = await updateAssignmentsWithSend(tokens[t]);
    }
    console.log('All users updated!');
    console.log('End.');
}

function textFromAssignmentIds(assignmentIds, rawAssignments) {
    let subjects = moodle.getAssignmentIdsSubjectNameByRawAssignments(rawAssignments);
    if (assignmentIds.length > 1) {
        let c = `Du hast in diesen Fächern neue Aufgaben:\n`;
        for (let a in assignmentIds) {
            let ex = `\n`;
            if (a == assignmentIds.length) {
                ex == '';
            }
            c = c + subjects[assignmentIds[a].toString()] + ex;
        }
        return c;
    } else {
        return `Du hast eine neue Aufgabe in ${subjects[assignmentIds[0].toString()]}`;
    }
}

async function stopServiceForAssignment(assignment, reason='') {
    console.log(`stopping service for acc with token "${assignment.moodleToken}".`);
    await sendSignalMessage(assignment.signalNumber, `Du wirst nicht länger Benachrichtigungen bekommen. ${reason}`);
    await assignment.destroy();
}


async function getAllAccountTokens() {
    let accs =  await db.Assignment.findAll({
        attributes: ['moodleToken']
    });
    let a = [];
    for (let ac in accs) {
        a.push(accs[ac].dataValues.moodleToken);
    }
    return a;
}


module.exports = {
    updateAssignmentsWithSend: updateAssignmentsWithSend,
    updateAllAssignmentsWithSend, updateAllAssignmentsWithSend,
    sendSignalMessage
    // updateAssignments: updateAssignments,
}