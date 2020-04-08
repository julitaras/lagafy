import React, { useState, useEffect } from "react";
import { userContext, userDispatch } from "../../context";
import {
  getTokenFromLocalStorage,
  setTokenInLocalStorage,
  isTokenExpired
} from "../../helpers/auth";
import { UserAgentApplication } from "msal";
import Welcome from "../Welcome";
import config from "../../configs/config";
import { useRouter } from "next/router";
import css from "../../../style.scss";

const Index = () => {
  const { user, token } = userContext();
  const dispatch = userDispatch();
  const [error, setError] = useState(false);
  const [loading, setLoading] = useState(false);
  const router = useRouter();
  let userAgentApplication;

  const login = async () => {
    return await userAgentApplication.loginRedirect({
      scopes: ["https://graph.microsoft.com/User.ReadWrite"]
    });
  };

  const getUserData = async () => {
    const userData = await userAgentApplication.getAccount();
    return JSON.parse(JSON.stringify(userData, null, 4));
  };

  const isUserLogged = token => {
    return token !== null && !isTokenExpired();
  };

  const getUserProfile = async token => {
    setLoading(true);

    try {
      var user = userAgentApplication.getAccount();
      dispatch({
        type: "LOGIN",
        user: user,
        token: token.idToken.rawIdToken
      });
      router.push("/home");
    } catch (err) {
      setError(true);
      setLoading(false);
    }
  };

  useEffect(() => {
    userAgentApplication = new UserAgentApplication({
      auth: {
        clientId: config.appId,
        redirectUri: config.redirectUri,
        authority: "https://login.microsoftonline.com/common",
        validateAuthority: true,
        //postLogoutRedirectUri: "https://localhost:3000/",
        navigateToLoginRequestUrl: false
      },
      cache: {
        cacheLocation: "localStorage",
        storeAuthStateInCookie: true
      }
    });

    userAgentApplication.handleRedirectCallback(async (err, token) => {
      if (err) setError(true);
      else {
        setTokenInLocalStorage(token);
        try {
          const userData = await getUserData();
          dispatch({
            type: "LOGIN",
            user: userData,
            token: token.idToken.rawIdToken
          });
          router.push("/home");
        } catch (err) {
          setError(true);
        }
      }
    });
  });

  useEffect(() => {
    const token = getTokenFromLocalStorage();
    if (isUserLogged(token)) {
      getUserProfile(token);
    }
  }, []);

  return (
    <>
      {loading ? (
        <div className={css.containerSpinner}>
          <div className="spinner-border text-primary" role="status">
            <span className="sr-only">Loading...</span>
          </div>
        </div>
      ) : (
        <>
          <div>
            <Welcome
              isAuthenticated={user !== null}
              user={user}
              authButtonMethod={() => login()}
            />
          </div>
        </>
      )}
    </>
  );
};

export default Index;
