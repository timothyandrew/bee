import logo from './logo.svg';
import './App.css';
import './wasm';
import React, { useState, useEffect } from 'react';

function App() {
  useEffect(() => {
    const go = new window.Go();
    WebAssembly.instantiateStreaming(fetch("bee.wasm"), go.importObject).then((result) => {
      go.argv = ["bee", "y", "nailgz"];
      go.run(result.instance);
  });

  });

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
