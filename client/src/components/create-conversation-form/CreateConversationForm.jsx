import React, { useState } from "react";
import style from "./CreateConversationForm.module.css";
import DefaultModal from "../../ui/default-modal/DefaultModal";
import { conversationAPI } from "../../apis/conversation.api";
import { useNavigate } from "react-router-dom";

const CreateConversationForm = ({ friends, handleClose }) => {
  const navigate = useNavigate()
  const [input, setInput] = useState("")
  const [users, setUsers] = useState([])
  const [friendsList, setFriendsList] = useState(friends)

  const filterFriends = ({ target }) => {
    setInput(target.value)
    const list = friends.filter(({ user }) => user.username.toLowerCase().includes(target.value.toLowerCase()))
    setFriendsList(list)
  }

  const updateUsers = ({ target }) => {
    const remove = users.includes(target.value)

    if (remove) {
      setUsers(users.filter(u => u !== target.value))
    } else {
      setUsers(users.concat(target.value))
    }
  }

  const onCreateConversation = async e => {
    e.preventDefault()
    const conversation = await conversationAPI.create(users)
    navigate(`/channels/@me/${conversation.id}`)
    handleClose()
  }

  return (
    <DefaultModal handleClose={handleClose}>
      <div className={style.container}>
        <span className={style.title}> Select Friends </span>
        <span className={style.subtitle}> You can add {9 - users.length} more friends </span>
        <form className={style.form} onSubmit={onCreateConversation}>
          <input className={style.input} onChange={filterFriends} value={input} placeholder="Type the username of a friend" />
          <div className={style.friends}>
            {
              friendsList.map(({ user }) => (
                <div key={user.id}>
                  <span> {user.username} </span>
                  <input type="checkbox" value={user.id} onChange={updateUsers} checked={users.includes(user.id)} />
                </div>
              ))
            }
          </div>
          <div className={style.separator} />
          <input type="submit" className={style.submit} value={users.length > 1 ? "Create Group DM" : "Create DM"} />
        </form>
      </div>
    </DefaultModal>
  )
}

export default CreateConversationForm