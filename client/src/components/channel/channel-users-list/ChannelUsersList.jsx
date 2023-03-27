import React, { useEffect, useState } from "react";
import User from "../../user/User";
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
              <p className={style.roleTitle}> {role.name.toUpperCase()} - {usersByRole[r].length} </p>
              {
                usersByRole[r].map(u => {
                  const { id, username, status, customStatus, avatar } = users.find(user => user.id === u)
                  return (
                    <User key={id} username={username} status={status} customStatus={customStatus} avatar={avatar} color={role.color} />
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