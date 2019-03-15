import * as React from "react";
import { PagesApi, LoginRequest, SigninRequest } from "../../api";

const minNameLength = 2
const maxNameLength = 32

const minPwdLength = 10
const maxPwdLength = 32

export interface PropsState {}
export interface SignonState { 
    form: {username: string, email: string, password: string},
    usernameValid: boolean,
    passwordValid: boolean,
    emailValid: boolean
}

export class Signon extends React.Component<PropsState, SignonState>{
    private api = new PagesApi()

    constructor(props: PropsState, state: SignonState){
        super(props, state)

        this.state = {
            form: {username: '', email:'', password: ''},
            usernameValid: true,
            passwordValid: true,
            emailValid: true
        } as SignonState
        
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    private set usernameValid(v: boolean){
        console.log("Valid", v)
        this.setState(() => {
            return {usernameValid: v}
        });
    }

    private set passwordValid(v: boolean){
        this.setState(() => {
            return {passwordValid: v}
        });
    }

    private validate(name: string){
        this.api.nameIsFree(name)
            .then(()=>{this.usernameValid = true})
            .catch(()=>{this.usernameValid = false})
    }

    private validatePassword(pwd: string){
        if(pwd.length < minPwdLength || pwd.length > maxPwdLength) { 
            this.passwordValid = false
        }else{
            this.passwordValid = true
        }
    }


    private handleChange(event: React.ChangeEvent<HTMLInputElement>) {
        const target = event.target;
        const value = target.value;
        const trigger = target.type === 'checkbox' ? target.checked: false
        const name = target.name;

        this.setState((state): SignonState => {
            let mystate;
            switch(name){
                case "username": 
                case "email":
                case "password":{
                    mystate = state.form
                    break
                }

                default:{
                    mystate = state
                }
            }

            mystate[name] = value

            return state
        });

        if (name === "username") {
            this.validate(value as string)
        } else if (name == "password"){
            this.validatePassword(value as string)
        }
    }
    
    private handleSubmit(event) {
        event.preventDefault();

        let form = this.state.form 
        console.log(form)

        this.api.signin(form as SigninRequest).then(resp => {
            let status = resp.status
            if(status !== 200){return}

            console.log("SIGNED ON!")
            window.open("/","_self")
        })
    }

    render() {
        let validNameClass = this.state.usernameValid ? "is-valid" : "is-invalid"
        let validPwdClass = this.state.passwordValid ? "is-valid" : "is-invalid"
        let validEmailClass = this.state.emailValid ? "is-valid" : "is-invalid"

        return (
            <div className="menu-block-wrap">
            <h1 className="text-center mb-4">Регистрация</h1>
            <form onSubmit={this.handleSubmit} className="menu-block form-signin needs-validation">
                <div className="form-label-group has-error has-feedback form-group">
                <input 
                            defaultValue={this.state.form.username} 
                            onChange={this.handleChange} 
                            type="username" id="username" name="username" 
                            autoCapitalize="none" 
                            autoCorrect="off" 
                            autoComplete="username" 
                            minLength={minNameLength} maxLength={maxNameLength}
                            className={"form-control " + validNameClass } 
                            placeholder="Имя пользователя" 
                            required={false}
                        />

                    <label htmlFor="username">Имя пользователя</label>
                    <div className="valid-feedback">
                        Looks good!
                    </div>
                    <div className="invalid-feedback">
                        Something is wrong!
                    </div>
                </div>

                <div className="form-label-group has-error has-feedback form-group">
                <input 
                            defaultValue={this.state.form.email} 
                            onChange={this.handleChange} 
                            type="text" id="email" name="email" 
                            autoCapitalize="none" 
                            autoCorrect="off" 
                            autoComplete="email" 
                            className={"form-control " + validEmailClass } 
                            placeholder="E-mail"
                            required={false}
                        />

                    <label htmlFor="email">E-mail</label>
                    <div className="valid-feedback">
                        Looks good!
                    </div>
                    <div className="invalid-feedback">
                        Something is wrong!
                    </div>
                </div>

                <div className="form-group">
                    <div className="form-label-group">
                        <input 
                            defaultValue={this.state.form.password} 
                            onChange={this.handleChange} 
                            type="password" 
                            id="inputPassword" 
                            name="password" 
                            className={"form-control " + validPwdClass}
                            placeholder="Password" 
                            autoCapitalize="none" 
                            autoCorrect="off"
                            minLength={minPwdLength} maxLength={maxPwdLength}
                            required={false} 
                        />
                        <label htmlFor="inputPassword">Пароль</label>
                    </div>
                    
                    <div className="form-label-group">
                    <input 
                            defaultValue={this.state.form.password} 
                            onChange={this.handleChange} 
                            type="password" 
                            id="repeatPassword"
                            className={"form-control " + validPwdClass}
                            placeholder="Repeat Password" 
                            autoCapitalize="none" 
                            autoCorrect="off" 
                            autoComplete="password" 
                            minLength={minPwdLength} maxLength={maxPwdLength}
                            required={false} 
                        />
                        <label htmlFor="repeatPassword">Повторите пароль</label>
                    </div>
                </div>

                <div className="checkbox mb-3">
                    <label>
                        <input type="checkbox" value="remember-me"/> Принимаю пользовательское соглашение
                    </label>
                </div>

                <button className="btn btn-lg btn-success btn-block" type="submit">Отправить</button>                                      
            </form>
            </div>
        );
    }
}
