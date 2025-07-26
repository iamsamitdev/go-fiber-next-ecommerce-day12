import React, { useState, useEffect } from "react"

interface Option {
  label: string
  value: string | number
}

interface CustomSelectProps {
  options: Option[]
}

const CustomSelect = ({ options }: CustomSelectProps) => {
  const [isOpen, setIsOpen] = useState(false)
  const [selectedOption, setSelectedOption] = useState<Option>(options[0])

  const toggleDropdown = () => {
    setIsOpen(!isOpen)
  }

  const handleOptionClick = (option: Option) => {
    setSelectedOption(option)
    toggleDropdown()
  }

  useEffect(() => {
    // closing modal while clicking outside
    function handleClickOutside(event: MouseEvent) {
      const target = event.target as Element
      if (!target.closest(".dropdown-content")) {
        toggleDropdown()
      }
    }

    if (isOpen) {
      document.addEventListener("mousedown", handleClickOutside)
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside)
    }
  }, [isOpen, toggleDropdown])

  return (
    <div className="dropdown-content custom-select relative min-w-[180px]">
      <div
        className={`select-selected whitespace-nowrap ${
          isOpen ? "select-arrow-active" : ""
        }`}
        onClick={toggleDropdown}
      >
        {selectedOption.label}
      </div>
      <div className={`select-items ${isOpen ? "" : "select-hide"}`}>
        {options.map((option: Option, index: number) => (
          <div
            key={index}
            onClick={() => handleOptionClick(option)}
            className={`select-item ${
              selectedOption === option ? "same-as-selected" : ""
            }`}
          >
            {option.label}
          </div>
        ))}
      </div>
    </div>
  )
}

export default CustomSelect