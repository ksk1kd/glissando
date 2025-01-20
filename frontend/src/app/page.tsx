'use client'
import { useEffect, useState } from 'react'

import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'

const views = [
  {
    id: 'monthly',
    label: 'Monthly View',
    list_name: 'months',
  },
  {
    id: 'member',
    label: 'Member View',
    list_name: 'members',
  },
  {
    id: 'project',
    label: 'Project View',
    list_name: 'projects',
  },
]
const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
const members = ['Alice', 'Bob', 'Charlie', 'David', 'Eve', 'Frank', 'Grace', 'Hank', 'Ivy', 'Jack']
const projects = Array.from({ length: 20 }, (_, i) => `Project ${i + 1}`)

export default function Page() {
  const [currentView, setCurrentView] = useState(views[0].id)
  const [displays, setDisplays] = useState(months)
  const [currentDisplay, setCurrentDisplay] = useState(displays[0])
  const [columns, setColumns] = useState(members)
  const [rows, setRows] = useState(projects)
  const [resources, setResources] = useState(Array.from({ length: rows.length }, () => Array(columns.length).fill(0)))

  useEffect(() => {
    const selectedView = views.find((view) => view.id === currentView)
    if (selectedView) {
      setDisplays(eval(selectedView.list_name))
      setCurrentDisplay(eval(selectedView.list_name)[0])
      let newRows: string[] = []
      let newColumns: string[] = []
      switch (selectedView.id) {
        case 'monthly':
          newRows = projects
          newColumns = members
          break
        case 'member':
          newRows = projects
          newColumns = months
          break
        case 'project':
          newRows = members
          newColumns = months
          break
      }
      setRows(newRows)
      setColumns(newColumns)
      setResources(Array.from({ length: newRows.length }, () => Array(newColumns.length).fill(0)))
    }
  }, [currentView])

  return (
    <>
      <main className="w-[1000px] mx-auto p-10">
        <header className="flex items-center gap-4 mb-5">
          <Select value={currentView} onValueChange={setCurrentView}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Select View" />
            </SelectTrigger>
            <SelectContent>
              {views.map((view, index) => (
                <SelectItem value={view.id} key={view.id}>
                  {view.label}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          <Select value={currentDisplay} onValueChange={setCurrentDisplay}>
            <SelectTrigger className="w-[180px]">
              <SelectValue placeholder="Select Display" />
            </SelectTrigger>
            <SelectContent>
              {displays.map((display, index) => (
                <SelectItem value={display} key={display}>
                  {display}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </header>
        <table className="table-auto w-full">
          <thead>
            <tr>
              <th></th>
              {columns.map((column, index) => (
                <th key={index} className="p-1">
                  {column}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {rows.map((row, rowIndex) => (
              <tr key={rowIndex}>
                <th className="p-1 w-40">{row}</th>
                {resources[rowIndex].map((value, colIndex) => (
                  <td key={`${rowIndex}-${colIndex}`} className="p-1">
                    <Input className="text-right" type="number" placeholder={value} />
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </main>
    </>
  )
}
