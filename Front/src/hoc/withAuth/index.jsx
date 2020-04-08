import React, { useState, useContext, useEffect } from "react";

import { useRouter } from "next/router";
import { userContext } from "../../context";
import { isUserLogged, isTokenExpired } from "../../helpers/auth";
import css from "../../.././style.scss";

const Index = WrappedComponent => () => {
  const { user } = userContext();
  const router = useRouter();

  useEffect(() => {
    if (!isUserLogged(user) || isTokenExpired()) {
      router.push("/");
    }
  }, []);

  return isUserLogged(user) ? (
    <WrappedComponent />
  ) : (
    <div className={css.containerSpinner}>
      <div className="spinner-border text-primary" role="status">
        <span className="sr-only">Loading...</span>
      </div>
    </div>
  );
};

export default Index;
