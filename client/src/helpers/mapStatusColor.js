const mapStatusColor = status => {
  const statusColor = {
    "Online": "#007000",
    "Idle": "#ceb900",
    "Do Not Disturb": "#8b0f0f",
    "Offline": "#777777",
  }

  return statusColor[status]
}

export default mapStatusColor