import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import DefaultModal from "../../ui/default-modal/DefaultModal";
import style from "./CreateServerForm.module.css";
import DefaultAlert from "../../ui/default-alert/DefaultAlert";
import { userAPI } from "../../apis/user.api";
import { createServer } from "../../store/slices/server.slice";
import { AiOutlineCloudUpload } from "react-icons/ai";

const CreateServerForm = ({ handleClose }) => {
  const dispatch = useDispatch()
  const { username } = useSelector(state => state.user)
  const [input, setInput] = useState({ name: `${username}'s server`, avatar: "" })

  const handleUpload = async ({ target }) => {
    try {
      const formData = new FormData()
      formData.append("file", target.files[0])
      const file = await userAPI.uploadAvatar(formData)
      setInput({ ...input, avatar: file })
    } catch (err) {
      DefaultAlert({
        icon: "error",
        title: "Image upload failed",
        text: err?.response?.data?.message || err?.code,
        timer: 3000
      })
    }
  }

  const onCreateServer = e => {
    e.preventDefault()
    dispatch(createServer(input))
      .then(() => handleClose())
  }

  useEffect(() => () => setInput({ name: `${username}'s server`, avatar: "" }), [username]);

  return (
    <DefaultModal handleClose={handleClose}>
      <div className={style.container}>
        <h3> Create a server </h3>
        <h4> Your server is where you and your friends hang out. Make yours and start talking. </h4>
        <form className={style.form} onSubmit={onCreateServer}>
          <div className={style.upload}>
            <label htmlFor="upload" className={style.uploadLabel}>
              {
                input.avatar
                  ? <img src={input.avatar && `${process.env.REACT_APP_API}/uploads/${input.avatar}`} alt="" className={style.image} />
                  : <AiOutlineCloudUpload size={32} className={style.uploadButton} />
              }
            </label>
            <input type="file" name="avatar" id="upload" hidden onChange={handleUpload} accept="image/png, image/jpeg" multiple={false} />
          </div>
          <div className={style.name}>
            <label> Server Name </label>
            <input value={input.name} className={style.input} onChange={e => setInput({ ...input, name: e.target.value })} />
          </div>
          <input type="submit" className={style.submit} value="Create" />
        </form>
      </div>
    </DefaultModal>
  )
}

export default CreateServerForm