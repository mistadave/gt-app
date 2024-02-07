/* Use the exported functions to call the API.
 * If necessary, adjust the backend address below:
 */
// const backend = "https://we3-finance-srv.sifs0005.infs.ch";
const backend = "http://localhost:8080/v1";

export function login(login, password) {
  return postJson("/auth/login", { login, password }).then(parseJSON);
}

export function signup(login, firstname, lastname, password) {
  return postJson("/auth/register", {
    login,
    firstname,
    lastname,
    password,
  }).then(parseJSON);
}

export function getGames() {
  return getJson(`/games`).then(parseJSON);
}

export function getGameLink(id) {
  return getJson(`/games/${id}/links`).then(parseJSON);
}

export function getAccount(accountNr, token) {
  return getAuthenticatedJson(`/accounts/${accountNr}`, token).then(parseJSON);
}

export function transfer(target, amount, token) {
  return postAuthenticatedJson("/accounts/transactions", token, {
    target,
    amount,
  }).then(parseJSON);
}

export function getTransactions(
  token,
  fromDate = "",
  toDate = "",
  count = 3,
  skip = 0
) {
  return getAuthenticatedJson(
    `/accounts/transactions?fromDate=${fromDate}&toDate=${toDate}&count=${count}&skip=${skip}`,
    token
  ).then(parseJSON);
}

function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response;
  } else {
    const error = new Error(response.statusText);
    error.response = response;
    throw error;
  }
}

function parseJSON(response) {
  return response.json();
}

function getAuthenticatedJson(endpoint, token) {
  return fetch(`${backend}${endpoint}`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
      Accept: "application/json",
    },
  }).then(checkStatus);
}

function getJson(endpoint) {
  return fetch(`${backend}${endpoint}`, {
    method: "GET",
    headers: {
      Accept: "application/json",
    },
  }).then(checkStatus);
}

function postJson(endpoint, params) {
  return fetch(`${backend}${endpoint}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify(params),
  }).then(checkStatus);
}

function postAuthenticatedJson(endpoint, token, params) {
  return fetch(`${backend}${endpoint}`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify(params),
  }).then(checkStatus);
}
