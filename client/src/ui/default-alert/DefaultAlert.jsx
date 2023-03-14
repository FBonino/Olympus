import Swal from "sweetalert2"

const DefaultAlert = ({ icon, title, text, confirmText = null, timer = null }) => {
  return Swal.fire({
    background: "#222831",
    color: "#DDDDDD",
    icon,
    title,
    text,
    timer,
    confirmButtonText: confirmText,
    confirmButtonColor: "#F05454"
  })
}

export default DefaultAlert