import React, { useState } from "react";

function App() {
  const [puuid, setPUUID] = useState(0);

  fetch("http://127.0.0.1:8080/summoner?name=doublelift")
    .then((res) => res.json())
    .then((result) => {
      setPUUID(result.puuid);
    });

  return <div>{puuid}</div>;
}

export default App;
