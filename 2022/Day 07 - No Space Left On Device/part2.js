// AoC 2022 Day 07 - No Space Left On Device - Part 1

const { readFileSync } = require('fs')

const availableSpace = 70000000
const minimumFreeSpace = 30000000
const maximumUsedSpace = availableSpace - minimumFreeSpace

const terminalOutput = readFileSync('input.txt', 'utf8').split('\n')

const cwdPath = []
let cwd

function makeDir (name, dir) {
  const newDir = { name, type: 'dir', size: 0, files: [] }
  if (dir !== null && dir !== undefined) dir.files.push(newDir)
  return newDir
}

function changeDir (path, dir) {
  let newDir = dir
  if (typeof path === 'string') {
    path = path.split('/')
  }
  path.forEach((dirName) => {
    newDir = newDir.files.find(d => d.type === 'dir' && d.name === dirName)
  })
  return newDir
}

function makeFile (name, size, dir) {
  const newFile = { name, type: 'file', size }
  dir.files.push(newFile)
  return newFile
}

function updateSize (dir) {
  const sizeofDirs = dir.files.filter((d) => d.type === 'dir').reduce((a, d) => a + updateSize(d), 0) ?? 0
  const sizeOfFiles = dir.files.filter((f) => f.type === 'file').reduce((a, f) => a + f.size, 0) ?? 0
  dir.size = sizeofDirs + sizeOfFiles
  return dir.size
}

const fileSystem = terminalOutput.reduce((fs, line) => {
  if (line.match(/^\$ cd /)) {
    const [,, dirName] = line.match(/^\$ (cd|ls) ?(.+)?/)
    if (dirName === '/') {
      cwd = fs
      cwdPath.length = 0
    } else if (dirName === '..') {
      cwdPath.pop()
      cwd = changeDir(cwdPath, fs)
    } else if (cwd.files.find((f) => f.type === 'dir' && f.name === dirName)) {
      cwd = changeDir(dirName, cwd)
      if (cwd === null) {
        throw new Error('changeDir() resulted in null working directory')
      }
      cwdPath.push(dirName)
    } else {
      cwd.files.push(cwd = makeDir(dirName, cwd))
      cwdPath.push(dirName)
    }
  } else if (line.match(/^\d+ /)) {
    const [, size, fileName] = line.match(/^(\d+) ([a-z.]+)/)
    makeFile(fileName, parseInt(size), cwd)
  } else if (line.match(/^dir/)) {
    const [, dirName] = line.match(/^dir ([a-z.]+)/)
    makeDir(dirName, cwd)
  }
  return fs
}, makeDir('/'))

updateSize(fileSystem)

function getDirs (dir, predicate = (d) => true) {
  const dirs = []
  dir.files.forEach((d) => {
    if (d.type === 'dir') {
      if (predicate(d)) dirs.push(d)
      dirs.push(...getDirs(d, predicate))
    }
  })
  return dirs
}

const currentUsedSpace = fileSystem.size

console.log(
  getDirs(fileSystem, (d) => currentUsedSpace - d.size < maximumUsedSpace)
    .map((d) => ({ name: d.name, type: d.type, size: d.size, spaceUsed: currentUsedSpace - d.size }))
    .sort((a, b) => -(a.spaceUsed - b.spaceUsed))
    // eslint-disable-next-line no-unexpected-multiline
    [0]
    .size
)
