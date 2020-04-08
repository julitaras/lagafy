import React from "react";
import css from "../../../style.scss";

function WelcomeContent(props) {
  // If authenticated, greet the user
  if (props.isAuthenticated) {
    return (
      // <div>
      //   <h4>Welcome {props.user.displayName}!</h4>
      //   <p>Use the navigation bar at the top of the page to get started.</p>
      // </div>
      null
    );
  }

  // Not authenticated, present a sign in button
  return (
    <div className="container">
      <div className="row col-center-text">
        <div className={css.formSignin}>
          <div className="text-center">
            <i className="icon-Logo-01"></i>
            <h1 className={css.titleLogin}>Lagafy</h1>
          </div>
          <button className="btn btnLogin" onClick={props.authButtonMethod}>
            Entrar
          </button>
        </div>
      </div>
    </div>
  );
}

export default class Welcome extends React.Component {
  render() {
    return (
      <WelcomeContent
        isAuthenticated={this.props.isAuthenticated}
        user={this.props.user}
        authButtonMethod={this.props.authButtonMethod}
      />
    );
  }
}
