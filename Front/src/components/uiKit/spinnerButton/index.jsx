import React from "react";
import css from "../../../../style.scss";
import cssSpinnerButton from "./index.scss";
import { isCheckinStarted } from "../../../helpers/dateHelper";

const index = ({
  children,
  loading,
  className,
  onClick,
  disabled,
  checkined
}) => {
  // return (
  //   <button className={className} onClick={onClick} disabled={disabled}>
  //     {true ? (
  //       <div>
  //         <span
  //           className={css.Spinner}
  //           class="spinner-border spinner-border-sm text-center"
  //           role="status"
  //           aria-hidden="true"
  //         ></span>
  //       </div>
  //     ) : (
  //       <div className={cssSpinnerButton.CheckinConfirmed}>
  //         {checkined && (
  //           <span id="checkmark" className="icon-Checkmark-00"></span>
  //         )}
  //         <p>{children}</p>
  //       </div>
  //     )}
  //   </button>
  // );

  return (
    <button className={className} onClick={onClick} disabled={disabled}>
      {loading ? (
        <div className={cssSpinnerButton.SpinnerContainer}>
          <span
            className={css.Spinner}
            class="spinner-border spinner-border-sm text-center"
            role="status"
            aria-hidden="true"
          ></span>
        </div>
      ) : (
        <div className={cssSpinnerButton.CheckinConfirmed}>
          {checkined && (
            <span id="checkmark" className="icon-Checkmark-00"></span>
          )}
          <p>{children}</p>
        </div>
      )}
    </button>
  );
};

export default index;
