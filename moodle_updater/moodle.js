const moodle_client = require('moodle-client');
const request_promise = require('request-promise');

module.exports = {
    getTokenByCredentials: getTokenByCredentials,
    getAssignmentIdsSubjectNameByRawAssignments: getAssignmentIdsSubjectNameByRawAssignments,
    getAssignmentIdsByRawAssignments: getAssignmentIdsByRawAssignments,
    getRawAssignmentsByToken: getRawAssignmentsByToken,
    getAssignmentIdsByToken: getAssignmentIdsByToken,
    validateToken: validateToken,
    checkResponseForInvalidToken
}

const MOODLE_HOST = process.env.MOODLE_HOST;
const MOODLE_TOKEN_SERVICE = process.env.MOODLE_TOKEN_SERVICE;

async function getTokenByCredentials(username, password) {
    var options = {
        uri: MOODLE_HOST + "/login/token.php",
        method: "POST",
        form: {
            service: MOODLE_TOKEN_SERVICE,
            username: username,
            password: password
        },
        strictSSL: true,
        json: true
    }

    return await request_promise(options)
        .then(function(res) {
            if ("token" in res) {
                return res.token;
            /*} else if ("error" in res) {
                return false; */
            } else {
                return false;
            }
        })
        .catch((err) => {
            console.log('Could not make a connection with moodle!')
        });
}


async function getClientByToken(token) {
    return await moodle_client.init({
        wwwroot: MOODLE_HOST,
        token: token
    });
}

async function getRawAssignmentsByToken(token) {
    let client = await getClientByToken(token);
    return await getRawAssignmentsByClient(client);
}

async function getAssignmentIdsByToken(token) {
    var client = await moodle_client.init({
        wwwroot: MOODLE_HOST,
        token: token
    });
    return await getAssignmentIdsByClient(client); 
}

async function getRawAssignmentsByClient(client) {
    return await client.call({
        wsfunction: "mod_assign_get_assignments",
        method: "POST"
    });
}

async function getAssignmentIdsByClient(client) {
    let raw = await getRawAssignmentsByClient(client);
    return getAssignmentIdsByRawAssignments(raw);
}

function getAssignmentIdsByRawAssignments(rawAssignments) {
    c = [];
    for (var a in rawAssignments['courses']) {
        for (var b in rawAssignments['courses'][a]['assignments']) {
            c.push(rawAssignments['courses'][a]['assignments'][b]['cmid']);
        }
    }
    return c;
}

function getAssignmentIdsSubjectNameByRawAssignments(rawAssignments) {
    c = {}
    for (var a in rawAssignments.courses) {
        for (var b in rawAssignments.courses[a].assignments) {
            c[rawAssignments.courses[a].assignments[b].cmid] = rawAssignments.courses[a].fullname;
        }
    }
    return c;
}

async function validateToken(token) {
    let client = await getClientByToken(token);
    let res = await client.call({
        wsfunction: 'enrol_self_get_instance_info',
        method: 'POST'
    });
    if (res.errorcode == 'invalidparameter') return true;
    return false;
}

function checkResponseForInvalidToken(response) {
    if (response.errorcode == 'invalidtoken') return true;
    return false;
}
