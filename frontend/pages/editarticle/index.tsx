import Header from '../../components/Header'
import Segment from '../../components/Segment'
import { Box, Button, FormLabel, Input, FormControl } from '@chakra-ui/react'
import dynamic from 'next/dynamic'
import { useEffect, useState } from 'react'
const CustomEditor = dynamic(() => import('../../components/EditorJS/Editor'), {
  ssr: false,
})

const EditArticle = () => {
  const [data, setData] = useState<any>()
  const [alias, setAlias] = useState('')
  const [title, setTitle] = useState('')

  const postArticle = () => {
    const headers = {
      accept: 'application/json',
      'Content-Type': 'application/json',
    }
    fetch('http://localhost:3000/api/articles', {
      method: 'POST',
      headers: headers,
      body: JSON.stringify({
        title,
        content: JSON.stringify(data),
        author_address: 'authorAddress',
        alias,
      }),
    }).then((response) => {
      if (!response.ok) {
        console.log('error', response.body)
      } else {
        console.log('ok', response.body)
      }
    })
  }

  useEffect(() => {
    console.log(JSON.stringify(data))
  })

  return (
    <>
      <Header></Header>
      <Segment>
        <Box>記事作成画面</Box>
      </Segment>
      <Segment>
        <Button onClick={postArticle}>投稿</Button>
        <FormControl>
          <FormLabel>Alias</FormLabel>
          <Input value={alias} onChange={(e: any) => setAlias(e.target.value)} />
        </FormControl>
        <FormControl>
          <FormLabel>Title</FormLabel>
          <Input value={title} onChange={(e: any) => setTitle(e.target.value)} />
        </FormControl>
        <Box bg='gray.500' w='100%' p={4}>
          <CustomEditor data={data} onChange={setData} holder='editorjs-container' />
        </Box>
      </Segment>
    </>
  )
}
export default EditArticle
