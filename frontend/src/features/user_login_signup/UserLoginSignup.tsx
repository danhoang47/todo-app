import { useState } from 'react';
import { faGoogle, faFacebook } from "@fortawesome/free-brands-svg-icons";

import { IconButton, Overlay } from "@/common/ui";

import "./UserLoginSignup.styles.scss";

function UserLoginSignup() {
	const [formType, setFormType] = useState<"login" | "signup">("login")

	return (
		<Overlay>
			<div id="userLoginSignupForm">
				<div className="modal">
					<div className="modalHeader">
						<h2 className="modalHeading">
							You must Sign In to Join
						</h2>
						<p className="modalSubHeading">
							We're A Team That Guides Each Other
						</p>
					</div>
					<div className="modalBody">
						<div className="socialLoginOptionList">
							<IconButton
								buttonClassName="socialLoginButton"
								icon={faGoogle}
							>
								<p>Signup with Google</p>
							</IconButton>
							<IconButton
								buttonClassName="socialLoginButton"
								icon={faFacebook}
							>
								<p>Signup with Facebook</p>
							</IconButton>
						</div>
						<div className="divider">Or</div>
						<form className="form loginForm">
							<div className="inputField userName">
								<label className="label" htmlFor="userName">Username</label>
								<input
									type="text"
									id="userName"
									placeholder="Username"
								/>
							</div>
							<div className="inputField password">
								<label className="label"  htmlFor="password">Password</label>
								<input
									type="text"
									id="password"
									placeholder="Password"
								/>
							</div>
                            <div className="checkBoxField rememberMeCheckBox">
                                <label className="label"  htmlFor="rememberMe">Remember Me</label>
								<input
									type="checkbox"
									id="rememberMe"
								/>
                            </div>
                            <button className="forgotPasswordBtn">Forgot Password?</button>
							<button className="submitBtn">Login</button>
						</form>
					</div>
				</div>
			</div>
		</Overlay>
	);
}

export default UserLoginSignup;
