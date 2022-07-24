export function getFileName (file: File): string {
  let key = Math.floor(Math.random() * 10).toFixed()
  key += Math.floor(new Date().getTime() / 1000).toFixed()
  for (let i = 0; i < 5; i++) {
    key += Math.floor(Math.random() * 10).toFixed()
  }
  key += getExt(file)
  return key
}

export  function getExt (file: File) {
  const dot = file.name.lastIndexOf('.')
  const ext = file.name.substring(dot)
  if (file.type === 'image/jpeg') {
      if (ext === '.jpg') {
      return '.jpg'
      }
  }
  if (file.type === 'image/png') {
      return '.png'
  }
  if (file.type === 'image/mp4') {
      return '.mp4'
  }
  return ext.toLowerCase()
}