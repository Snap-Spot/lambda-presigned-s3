# presigned-go
```js
import { useEffect, useRef, useCallback, useState } from "react";
import axios from "axios";

function App() {
  const inputRef = useRef;
  const [img, setImg] = useState("");
  const [imgMeta, setImgMeta] = useState();
  const [presignedUrl, setPresignedUrl] = useState("");
  const getPresignedUrl = () => {
    // console.log(img);
    // console.log(imgMeta);

    const a = axios
      .post(
        "url",
        {
          fileName: imgMeta.name,
        }
      )
      .then((res) => {
        const reader = new FileReader();
        axios.put(res.data, "파일", {
          "Content-Type": imgMeta.type,
        });
      });
  };

  const uploadImageInput = (event) => {
    var reader = new FileReader();
    reader.onload = function (event) {
      setImg(event.target.result);
    };
    setImgMeta(event.target.files[0]);
    reader.readAsDataURL(event.target.files[0]);
  };

  const onUploadImageButtonClick = useCallback(() => {
    if (!inputRef.current) {
      return;
    }
    inputRef.current.click();
  }, []);
  return (
    <div className="App">
      <h2>Hello world</h2>
      <input type="file" accept="image/*" onChange={uploadImageInput} />
      <button label="이미지 업로드" onClick={getPresignedUrl} />
    </div>
  );
}

export default App;

```
