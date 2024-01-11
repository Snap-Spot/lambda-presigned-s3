# presigned-go
## Summary
- 스냅스팟 웹서비스의 안전한 사진 업로드와 메인 서버의 업로드 딜레이를 줄이기 위해 AWS의 Lambda Function으로 S3 Presigned URL을 발급하고, API Gateway를 통해 호출합니다.
- 빠른 속도 보장을 위해 Go를 활용하였습니다.

## Architecture
<div align=center>
  <img width="441" alt="image" src="https://github.com/Snap-Spot/lambda-presigned-s3/assets/80109963/5f261a3c-17ab-4676-b787-8f2d2c32b3ba">
</div>

## Reeact Demo
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
