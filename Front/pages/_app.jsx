import React from "react";
/* eslint-disable */
import Head from "next/head";
import PropTypes from "prop-types";
import { UserProvider } from "../src/context";

const Index = ({ Component, pageProps }) => {
  return (
    <>
      <Head>
        <link
          rel="stylesheet"
          href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css"
          integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh"
          crossOrigin="anonymous"
        />
        <link rel="stylesheet" href="/css/fonts.css" />
      </Head>
      <UserProvider>
        <Component {...pageProps} />
      </UserProvider>
    </>
  );
};

export default Index;
