const graph = require("@microsoft/microsoft-graph-client");

function getAuthenticatedClient(accessToken) {
  // Initialize Graph client
  const client = graph.Client.init({
    // Use the provided access token to authenticate
    // requests
    authProvider: done => {
      done(null, accessToken.accessToken);
    }
  });

  return client;
}
/* eslint-disable */
export async function getUserDetails(accessToken) {
  /* eslint-enable */
  const client = getAuthenticatedClient(accessToken);

  const user = await client.api("/me").get();
  return user;
}
