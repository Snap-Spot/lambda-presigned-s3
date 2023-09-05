# presigned-go
```js
import logo from "./logo.svg";
import "./App.css";
import { useEffect, useRef, useCallback, useState } from "react";
import axios from "axios";

function App() {
  const inputRef = useRef(null);
  const [img, setImg] = useState("");
  const [imgMeta, setImgMeta] = useState();
  const [presignedUrl, setPresignedUrl] = useState("");
  const getPresignedUrl = () => {
    console.log(img);
    console.log(imgMeta);

    console.log(inputRef);

    const a = axios
      .post(
        "",
        {
          fileName: imgMeta.name,
        }
      )
      .then((res) => {
        axios.put(res.data, img, {
          "Content-Type": imgMeta.type,
        });
      });
  };

  const uploadImageInput = (event) => {
    setImgMeta(event.target.files[0]);
    var reader = new FileReader();
    /*
    reader.onload = function (event) {
      setImg(event.target.result);
    };
    */
    reader.readAsDataURL(imgMeta);
    reader.onload = () => {
      setImg(reader.result);
    };
  };

  return (
    <div className="App">
      <h2>Hello world</h2>
      <input
        type="file"
        accept="image/*"
        ref={inputRef}
        onChange={uploadImageInput}
      />
      <button label="이미지 업로드" onClick={getPresignedUrl} />
      <img width={"100%"} src={img} />
    </div>
  );
}

export default App;


```
