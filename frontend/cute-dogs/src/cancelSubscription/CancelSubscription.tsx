import React, {useState} from "react";
import env from "react-dotenv";
import "./CancelSubscription.css";

function CancelSubscription() {
    interface FormDataType {email:string}
    const formData: FormDataType = {email: ""}
    const [responseBody, setResponseBody] = useState<FormDataType>(formData)
    const [result, setResult] = useState<string>("")

    const inputChangeHandler = (event: React.ChangeEvent<HTMLInputElement>) => {
        const {name, value} = event.target
        setResponseBody({...responseBody, [name]:value})
      }

    const onSubmitHandler = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (responseBody.email === "") {
    setResult("O campo email não pode ser vazio")
    } else {
    await fetchDelete()
    }
    
    setResponseBody({...responseBody, "email":""})
    }

    async function fetchDelete() {
        let response = await fetch(`${env.REACT_APP_YOUR_DOMAIN}/api`, {
            method: 'DELETE',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
            email: responseBody.email,
  }),
  
    });
  let data = await response.text()

  setResult(data.replace(/['"]+/g, ''));
  }
    
    return(
        <div className="delete-container">
              <h2>Poxa, é uma pena que não queira mais receber doguinhos! :(</h2>
              <p>Digite abaixo o email que realizou a sua inscrição</p>
              <form onSubmit={async (e) => onSubmitHandler(e)}>
              <input type="email" name="email" value={responseBody.email}
              id="name-input"
              onChange={(e) => inputChangeHandler(e)}/>
              <div id="subscribe">
                <input type="submit" value="Enviar" />
              </div>
              </form>

            {result !== "" ?
            <div>
              <h4>{result}</h4>
            </div>: 
            <></>}
        </div>
        
    );
}

export default CancelSubscription;
