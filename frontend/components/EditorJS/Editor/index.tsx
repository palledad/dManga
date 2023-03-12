import React, { memo, useEffect, useRef } from 'react'
import EditorJS from '@editorjs/editorjs'
import Tools from '../Tools'
import { Box } from '@chakra-ui/react'

type Props = {
  data: any
  onChange: any
  holder: any
}

const CustomEditor: React.FC<Props> = ({ data, onChange, holder }) => {
  const ref = useRef<EditorJS | null>(null)

  useEffect(() => {
    if (!ref.current) {
      ref.current = new EditorJS({
        holder: holder,
        tools: Tools,
        autofocus: true,
        data,
        placeholder: 'Start writing your story...',
        async onChange(api) {
          const data = await api.saver.save()
          onChange(data)
        },
      })
    }
    //add a return function handle cleanup
    return () => {
      if (ref.current && ref.current.destroy) {
        ref.current.destroy()
      }
    }
  }, [])

  return <Box bg='white' id={holder}></Box>
}

// 状態の変更をこのFCの変更によってのみ検知する。
export default memo(CustomEditor)
