import React, {useState} from 'react';
import env from "react-dotenv";
import "./App.css"

function App() {
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
    await fetchCreate()
  }
  
  setResponseBody({...responseBody, ["email"]:""})
  }

  async function fetchCreate() {
    let response = await fetch(`${env.REACT_APP_YOUR_DOMAIN}/api`, {
      method: 'POST',
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

  return (
    <div className="App">
        <div>
          <h2>Insira seu email para receber doguinhos fofos todas as manhãs!</h2>
        </div>
        <div>
          <form onSubmit={async (e) => onSubmitHandler(e)}>
            <input type="email" name="email" value={responseBody.email}
            id="name-input"
            onChange={(e) => inputChangeHandler(e)}/>
            <div id="subscribe">
              <input type="submit" value="Inscrever-se" />
            </div>
          </form>
        </div>
        {result !== "" ?
        <div>
          <h4>{result}</h4>
        </div>: 
        <></>}
    </div>
  );
}

export default App;
