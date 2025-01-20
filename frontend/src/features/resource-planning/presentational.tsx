'use client'

import { useEffect, useState } from 'react'

import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { AxisType } from '@/types/axis'

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

export function ResourcePlanningPresentation({
  members,
  projects,
  months,
}: {
  members: AxisType[]
  projects: AxisType[]
  months: AxisType[]
}) {
  const [currentView, setCurrentView] = useState(views[0].id)
  const [displays, setDisplays] = useState(months)
  const [currentDisplay, setCurrentDisplay] = useState(displays[0].id)
  const [columns, setColumns] = useState(members)
  const [rows, setRows] = useState(projects)
  const [resources, setResources] = useState(Array.from({ length: rows.length }, () => Array(columns.length).fill(0)))
  const [columnSums, setColumnSums] = useState(
    columns.map((_, colIndex) => resources.reduce((sum, row) => sum + row[colIndex], 0)),
  )
  const [rowSums, setRowSums] = useState(resources.map((row) => row.reduce((sum, value) => sum + value, 0)))

  useEffect(() => {
    const selectedView = views.find((view) => view.id === currentView)
    if (selectedView) {
      setDisplays(eval(selectedView.list_name))
      setCurrentDisplay(eval(selectedView.list_name)[0].id)
      let newRows: AxisType[] = []
      let newColumns: AxisType[] = []
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
      const newResources = Array.from({ length: newRows.length }, () => Array(newColumns.length).fill(0))
      setResources(newResources)
      setColumnSums(newColumns.map((_, colIndex) => newResources.reduce((sum, row) => sum + row[colIndex], 0)))
      setRowSums(newResources.map((row) => row.reduce((sum, value) => sum + value, 0)))
    }
  }, [currentView, members, projects, months])

  const handleInputChange = (rowIndex: number, colIndex: number, value: number) => {
    const newResources = [...resources]
    newResources[rowIndex][colIndex] = value || 0
    setResources(newResources)
    setColumnSums(columns.map((_, colIndex) => newResources.reduce((sum, row) => sum + row[colIndex], 0)))
    setRowSums(newResources.map((row) => row.reduce((sum, value) => sum + value, 0)))
  }

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
              {displays.map((display) => (
                <SelectItem value={display.id} key={display.id}>
                  {display.name}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </header>
        <table className="table-auto w-full">
          <thead>
            <tr>
              <th></th>
              {columns.map((column) => (
                <th key={column.id} className="p-1 w-20">
                  {column.name}
                </th>
              ))}
              <th>Total</th>
            </tr>
          </thead>
          <tbody>
            {rows.map((row, rowIndex) => (
              <tr key={row.id}>
                <th className="p-1 w-40">{row.name}</th>
                {resources[rowIndex].map((value, colIndex) => (
                  <td key={`${rowIndex}-${colIndex}`} className="p-1">
                    <Input
                      className="text-right"
                      type="number"
                      value={value}
                      onChange={(e) => handleInputChange(rowIndex, colIndex, parseInt(e.target.value))}
                    />
                  </td>
                ))}
                <td className="p-1 text-right">{rowSums[rowIndex]}</td>
              </tr>
            ))}
            <tr>
              <th className="p-1 w-40">Total</th>
              {columnSums.map((sum, index) => (
                <td key={index} className="py-1 px-3 text-right">
                  {sum}
                </td>
              ))}
              <td className="p-1 text-right"></td>
            </tr>
          </tbody>
        </table>
      </main>
    </>
  )
}
