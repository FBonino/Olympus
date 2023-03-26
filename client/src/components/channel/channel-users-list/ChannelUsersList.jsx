import React, { useEffect, useState } from "react";
import style from "./ChannelUsersList.module.css";

const ChannelUsersList = ({ users, roles }) => {
  const [usersByRole, setUsersByRole] = useState()

  useEffect(() => {
    setUsersByRole(state => {
      const cache = Object.fromEntries(roles.map(r => [r.id, []]))
      for (const user of users) {
        const role = user.roles[0]
        cache[role] = cache[role].concat(user.id)
      }
      return cache
    })
  }, [users, roles])

  return (
    <div className={style.container}>
      {
        usersByRole && Object.keys(usersByRole).map(r => {
          const role = roles.find(role => role.id === r)
          return (
            <div className={style.role} key={r}>
              <p className={style.roleTitle}> {role.name} - {usersByRole[r].length} </p>
              {
                usersByRole[r].map(u => {
                  const user = users.find(user => user.id === u)
                  return (
                    <div className={style.user} key={u}>
                      <img className={style.avatar} src={`${process.env.REACT_APP_API}/uploads/${user.avatar}`} alt="" />
                      <div className={style.text}>
                        <span className={style.username} style={{ color: role.color }}> {user.username} </span>
                        <span className={style.customStatus}> {user.customStatus} </span>
                      </div>
                    </div>
                  )
                })
              }
            </div>
          )
        })
      }
    </div>
  )
}

export default ChannelUsersList