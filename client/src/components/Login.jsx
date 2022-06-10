import React from 'react'
import { useNavigate } from 'react-router'

const Login = () => {
  const navigate = useNavigate();

  const log = async (e) => {
    e.preventDefault();

    
    document.getElementById('error').innerHTML = "Вход в сервис...";
    const login_ = document.getElementById('login').value;
    const x = document.getElementById("myPsw").value;

    const data = { "login" : login_, "password" : x };
    console.log(data)
    fetch(`http://localhost:8080/login`, {
      method: 'POST', 
      body: JSON.stringify(data)
    }).then(response => response.json())
    .then(json => {
        console.log(json)
        if (json.Access == true) {
            localStorage.setItem(`user_id`, json.id);
            localStorage.setItem(`login`, json.login);
            navigate('/user/');
        }
        else {
            document.getElementById('error').innerHTML = "Неверный логин или пароль";
        }
    })
  };



  return (
    <div>
      <h1 className="mb-3">ooo"ПовезлоПовезло"</h1>
        <div>
          <p id="error"></p>
          <div className="input-group col-xs-3">                    
            <div className="input-group-append">
              <span>Логин</span>
              <input id="login" type="text" className="form-control" aria-describedby="basic-addon2" required />
              <br/>
              <br/>
              <br/>
              <span>Пароль</span>
              <input id="myPsw" type="password" className="form-control" aria-describedby="basic-addon2" required/>
            </div>
          </div>
          <br/>
          <button className="btn btn-primary" type="button" onClick={log}>Войти</button>
        </div>
    </div>
  )
}

export default Login