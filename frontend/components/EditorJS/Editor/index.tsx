import React, { memo, useEffect, useRef } from 'react'
import EditorJS from '@editorjs/editorjs'
import Tools from '../Tools'
import { Box } from '@chakra-ui/react'

// @ts-ignore
const CustomEditor = ({ data, onChange, holder }) => {
  //initialize editorjs
  const ref = useRef()

  useEffect(() => {
    //initialize editor if we don't have a reference
    if (!ref.current) {
      // @ts-ignore
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
      // @ts-ignore
      if (ref.current && ref.current.destroy) {
        // @ts-ignore
        ref.current.destroy()
      }
    }
  }, [])

  return <Box bg='white' id={holder}></Box>
}

// 状態の変更をこのFCの変更によってのみ検知する。
export default memo(CustomEditor)
