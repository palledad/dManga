import { Image, Box, Button } from '@chakra-ui/react'
import Link from 'next/link'
const Header = () => {
  return (
    <Box display='flex' alignItems='center' justifyContent='space-between' h={45} mx={4}>
      <Link href='/'>
        <Box display='flex'>
          <Image marginLeft={1} borderRadius='full' boxSize='30px' src='palledad_logo_1.jpg' />
        </Box>
      </Link>
      <Box display='flex' gap={4}>
        <Link href='/articlesearch'>
          <Button size='sm'>記事検索</Button>
        </Link>
        <Link href='/createarticle'>
          <Button size='sm'>記事作成</Button>
        </Link>
        <Link href='/mypage'>
          <Button size='sm'>アカウント</Button>
        </Link>
      </Box>
    </Box>
  )
}

export default Header
