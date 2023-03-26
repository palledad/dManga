import { List, ListItem } from '@chakra-ui/react'
import parse from 'html-react-parser'
import React from 'react'

const ListElement = (articleDic: any) => {
  return (
    <List as='ol' styleType='decimal'>
      {articleDic.data?.items.map((item: any, i: number) => {
        return <ListItem key={i}>{parse(item)}</ListItem>
      })}
    </List>
  )
}
export default ListElement
