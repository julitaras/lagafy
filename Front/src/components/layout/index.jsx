import Navbar from "../navbar";
import css from "../../../style.scss";
import { userContext, userDispatch } from "../../context";
import CONSTANTS from "../../configs/constants";

const Layout = props => {
  const { user, token } = userContext();

  const verifyUser = () => {
    return CONSTANTS.ADMINS.includes(user.userName);
  };

  return (
    <div className={css.containerIndex}>
      <div className={css.Layout}>
        <Navbar isAdmin={verifyUser()} currentPage={props.currentPage} />
        {props.children}
      </div>
    </div>
  );
};

export default Layout;
