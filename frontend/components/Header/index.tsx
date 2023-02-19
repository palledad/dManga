import { Image, Box, Button } from '@chakra-ui/react'
const Header = () => {
  return (
    <Box display='flex' alignItems='center' justifyContent='space-between' h={45} mx={4}>
      <Box display='flex'>
        <Image marginLeft={1} borderRadius='full' boxSize='30px' src='palledad_logo_1.jpg' />
      </Box>
      <Box display='flex' gap={4}>
        <Button size='sm'>記事検索</Button>
        <Button size='sm'>記事作成</Button>
        <Button size='sm'>アカウント</Button>
      </Box>
    </Box>
  )
}

export default Header
