import React from 'react';
import { 
  List, 
  ListItem, 
  ListItemText, 
  IconButton, 
  Typography,
  Box 
} from '@mui/material';
import { Add as AddIcon, Remove as RemoveIcon, Delete as DeleteIcon } from '@mui/icons-material';
import { useCart } from '../contexts/CartContext';

const CartItemList = () => {
  const { cart, updateQuantity, removeFromCart } = useCart();

  return (
    <List>
      {cart.map((item) => (
        <ListItem
          key={item.id}
          secondaryAction={
            <IconButton onClick={() => removeFromCart(item.id)}>
              <DeleteIcon />
            </IconButton>
          }
        >
          <ListItemText
            primary={item.name}
            secondary={`$${item.price.toFixed(2)} each`}
          />
          <Box sx={{ display: 'flex', alignItems: 'center', ml: 2 }}>
            <IconButton 
              onClick={() => updateQuantity(item.id, item.quantity - 1)}
              size="small"
            >
              <RemoveIcon />
            </IconButton>
            <Typography sx={{ mx: 2 }}>{item.quantity}</Typography>
            <IconButton 
              onClick={() => updateQuantity(item.id, item.quantity + 1)}
              size="small"
            >
              <AddIcon />
            </IconButton>
          </Box>
        </ListItem>
      ))}
    </List>
  );
};

export default CartItemList; 