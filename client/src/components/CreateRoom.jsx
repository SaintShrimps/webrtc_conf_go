import React from 'react';
import { useNavigate } from 'react-router';




const CreateRoom = () => {
    const navigate = useNavigate();

    const join = async (e) => {
        e.preventDefault();

        const room_id = document.getElementById('link').value;

        const data = { "RoomID" : room_id};
        console.log(data)
        fetch(`http://localhost:8080/GETInfRooms`, {
          method: 'POST', 
          body: JSON.stringify(data)
        }).then(response => response.json())
        .then(json => {
            console.log(json)
            if (json.Access == true) {
                navigate('/room/'+`${room_id}`);
            }
            else {
                document.getElementById('error').innerHTML = "Комнаты не существует";
            }
        })

    };
    
    const create = async (e) => {
        e.preventDefault();

        const resp = await fetch("http://localhost:8080/create");
        const { room_id } = await resp.json();

        navigate('/room/'+`${room_id}`);
    };

    return (
        <div>
            <h1 className="mb-3">ooo"ПовезлоПовезло"</h1>
            <div>
                <p id="error"></p>
                <div className="input-group col-xs-3">                    
                    <div className="input-group-append">
                    <span>Введите ID комнаты</span>
                    <input id="link" type="text" className="form-control" aria-describedby="basic-addon2" required /> 
                    <br/>
                    </div>
                </div>
                <br/>
                <button className="btn btn-primary" type="button" onClick={join}>Войти в комнату</button>
            
                <br /> 
                <br /> 
                <button className="btn btn-primary" type="button" onClick={create}>Создать комнату</button>

            </div>

        </div>
    );
};

export default CreateRoom;

