import logo from './logo.svg';
import './App.css';
import './wasm';
import React, { useState, useEffect } from 'react';

const DISPLAY_COUNT = 100;

function App() {
  useEffect(() => {
    const go = new window.Go();

    WebAssembly.instantiateStreaming(fetch("bee.wasm"), go.importObject).then((result) => {
      go.run(result.instance);
      setLoading(false);
    });

  }, []);

  const handleNonGoldChange = (e, i) => {
    setNonGold((ng) => {
      let copy = [...ng];
      copy[i] = e.target.value.trim();
      console.log(copy);
      return copy;
    });
  }

  let [loading, setLoading] = useState(true);

  let [gold, setGold] = useState(null);
  let [nonGold, setNonGold] = useState([]);

  if (loading) {
    return <div>Loading...</div>;
  }


  let results = JSON.parse(window.wasm_entry(gold, nonGold.filter(n => n).join(""))).result || [];
  let trimmedResults = (results.length > DISPLAY_COUNT) ? results.slice(DISPLAY_COUNT) : results;
  let hiddenCount = Math.max(results.length - DISPLAY_COUNT, 0);

  return (
    <div className="App mx-auto container px-4">
      <div className="flex space-x-12 justify-center mt-4">
        <div>
          <div>Gold</div>
          <input onChange={(e) => setGold(e.target.value)} className="border rounded border-slate-400 w-6 text-center"></input>
        </div>

        <div>
          <div>Non-Gold</div>
          <div className="flex space-x-4">
            <input onChange={(e) => handleNonGoldChange(e, 0)} className="border rounded border-slate-400 w-6 text-center"></input>
            <input onChange={(e) => handleNonGoldChange(e, 1)} className="border rounded border-slate-400 w-6 text-center"></input>
            <input onChange={(e) => handleNonGoldChange(e, 2)} className="border rounded border-slate-400 w-6 text-center"></input>
            <input onChange={(e) => handleNonGoldChange(e, 3)} className="border rounded border-slate-400 w-6 text-center"></input>
            <input onChange={(e) => handleNonGoldChange(e, 4)} className="border rounded border-slate-400 w-6 text-center"></input>
            <input onChange={(e) => handleNonGoldChange(e, 5)} className="border rounded border-slate-400 w-6 text-center"></input>
          </div>
        </div>
      </div>

      {results && <ul className="mt-4">
        {trimmedResults.map((r, i) => {
          return <li key={i}>
            {r}
          </li>
        })}
        {hiddenCount > 0 && <li>and {hiddenCount} more</li>}
      </ul>}
    </div>
  );
}

export default App;
