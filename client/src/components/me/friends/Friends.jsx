import React, { useEffect, useState } from "react";
import FriendsNav from "../friends-nav/FriendsNav";
import style from "./Friends.module.css";
import { useSelector } from "react-redux";
import Friend from "../friend/Friend";
import { RxMagnifyingGlass } from "react-icons/rx";

const Friends = () => {
  const [search, setSearch] = useState("")
  const [selected, setSelected] = useState("Online")
  const { friends } = useSelector(state => state.user)
  const [friendsList, setFriendsList] = useState(friends)

  useEffect(() => {
    let list
    if (selected === "All friends") {
      list = friends.filter(f => f.relation === "Friend")
    } else {
      list = friends.filter(f => f.relation === selected)
    }
    list = list.filter(f => f.user.username.toLowerCase().includes(search.toLowerCase()))
    setFriendsList(list)
  }, [selected, friends, search])

  return (
    <div className={style.container}>
      <div className={style.nav}>
        <FriendsNav selected={selected} setSelected={setSelected} resetSearch={() => setSearch("")} />
      </div>
      <div className={style.content}>
        <div className={style.friends}>
          {
            (!!friendsList.length || search !== "") &&
            <>
              <div className={style.searchbar}>
                <input id="friendsInput" className={style.input} placeholder="Search" value={search} onChange={({ target }) => setSearch(target.value)} />
                <label htmlFor="friendsInput" className={style.label}> <RxMagnifyingGlass size={22} /> </label>
              </div>
              <div className={style.count}> {selected.toUpperCase()} - {friendsList.length} </div>
              <div className={style.friendsList}>
                {
                  friendsList.map(({ user }) => (
                    <Friend key={user.id} friend={user} />
                  ))
                }
              </div>
            </>
          }
        </div>
        <div className={style.activity}>
          <span className={style.title}> Active Now </span>
          <span className={style.subtitle}> It's quiet for now... </span>
          <span className={style.text}> When a friend starts an activity - like playing a game or hanging out on voice - we'll show it here! </span>
        </div>
      </div>
    </div>
  )
}

export default Friends